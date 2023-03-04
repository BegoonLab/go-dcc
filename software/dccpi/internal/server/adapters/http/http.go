package http

import (
	"context"
	"net/http"
	"time"

	"github.com/alexbegoon/go-dcc/internal/controller"

	"github.com/alexbegoon/go-dcc/internal/server/adapters/http/handlers"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

const (
	ShutdownTimeout = 3 * time.Second
	Timeout         = 3 * time.Second
	Address         = ":3000"
)

type Server struct {
	HTTPServer *http.Server
	logger     *zap.Logger
	quit       *atomic.Bool
	ctrl       *controller.Controller
}

func New(l *zap.Logger, ctrl *controller.Controller) *Server {
	server := &Server{
		logger: l,
		quit:   atomic.NewBool(false),
		ctrl:   ctrl,
	}

	h := server.registerHandlers()

	server.HTTPServer = &http.Server{
		Addr:         Address,
		Handler:      h,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}

	return server
}

func (s *Server) registerHandlers() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "ws://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/*", handlers.NewFileHandler(s.logger).ServeFiles)
	r.Get("/ws", handlers.NewWsHandler(s.logger, s.ctrl).ServeWS)
	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			s.logger.Error("Unable to respond", zap.Error(err))
		}
	})

	return r
}

func (s *Server) Serve() {
	s.logger.Info(spew.Sprintf("HTTP server is starting to listen at `%s` ...", Address))
	go func() {
		err := s.HTTPServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Error("Unable to start HTTP server", zap.Error(err))
		}
	}()
}

func (s *Server) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if s.quit.CAS(false, true) {
		if err := s.HTTPServer.Shutdown(ctx); err != nil {
			s.logger.Warn("HTTP server has been stopped by force", zap.Error(err))
		} else {
			s.logger.Info("HTTP server has been gracefully stopped")
		}
		s.quit.CAS(true, false)
	}
}
