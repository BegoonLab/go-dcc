package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type FileHandler struct {
	Logger *zap.Logger
}

func NewFileHandler(logger *zap.Logger) *FileHandler {
	return &FileHandler{
		Logger: logger,
	}
}

func (h *FileHandler) ServeFiles(w http.ResponseWriter, r *http.Request) {
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir("./../frontend/public")))
	fs.ServeHTTP(w, r)
}
