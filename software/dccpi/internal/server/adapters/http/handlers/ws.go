package handlers

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/alexbegoon/go-dcc/internal/controller"
	ctrlProto "github.com/alexbegoon/go-dcc/internal/pb/build/go/controller"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"
)

const (
	MaxMessagesInQueue = 16
	Burst              = 8
	Limit              = time.Millisecond * 100
)

// subscriber represents a subscriber.
// Messages are sent on the msgs channel and if the client
// cannot keep up with the messages, closeSlow is called.
type subscriber struct {
	id        string
	msgs      chan []byte
	closeSlow func()
}

type WsHandler struct {
	Logger *zap.Logger

	ctrl *controller.Controller

	// subscriberMessageBuffer controls the max number
	// of messages that can be queued for a subscriber
	// before it is kicked.
	//
	// Defaults to 16.
	subscriberMessageBuffer int

	// publishLimiter controls the rate limit applied to the publish endpoint.
	//
	// Defaults to one publish every 100ms with a burst of 8.
	publishLimiter *rate.Limiter

	subscribersMu sync.Mutex
	subscribers   map[*subscriber]struct{}
}

func NewWsHandler(logger *zap.Logger, ctrl *controller.Controller) *WsHandler {
	return &WsHandler{
		Logger:                  logger,
		ctrl:                    ctrl,
		subscriberMessageBuffer: MaxMessagesInQueue,
		publishLimiter:          rate.NewLimiter(rate.Every(Limit), Burst),
		subscribers:             make(map[*subscriber]struct{}),
	}
}

func (h *WsHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	h.subscribeHandler(w, r)
}

func (h *WsHandler) readHandler(ctx context.Context, conn net.Conn) {
	go func() {
		for {
			var cProto ctrlProto.Controller
			msg, op, err := wsutil.ReadClientData(conn)
			if op != ws.OpBinary {
				h.Logger.Error(fmt.Sprintf("Client data is not binary, but %v", op))

				return
			}
			if err != nil {
				h.Logger.Error("Unable to read client data", zap.Error(err))

				return
			}
			err = proto.Unmarshal(msg, &cProto)
			if err != nil {
				h.Logger.Error("Unable to unmarshal client data", zap.Error(err))

				return
			}

			err = h.ctrl.Handle(&cProto)
			if err != nil {
				h.Logger.Error("Unable to Handle command from subscriber", zap.Error(err))

				return
			}

			go h.publish(ctx)
		}
	}()
}

// subscribeHandler accepts the WebSocket connection and then subscribes
// it to all future messages.
func (h *WsHandler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		h.Logger.Error("Unable to accept handshake", zap.Error(err))

		return
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			h.Logger.Error("Unable to close connection", zap.Error(err))
		}
	}(conn)

	err = h.subscribe(r.Context(), conn)
	if errors.Is(err, context.Canceled) {
		return
	}
	if err != nil {
		h.Logger.Error("Unable to subscribe", zap.Error(err))

		return
	}
}

// subscribe subscribes the given WebSocket to all broadcast messages.
// It creates a subscriber with a buffered msgs chan to give some room to slower
// connections and then registers the subscriber. It then listens for all messages
// and writes them to the WebSocket. If the context is canceled or
// an error occurs, it returns and deletes the subscription.
//
// It uses CloseRead to keep reading from the connection to process control
// messages and cancel the context if the connection drops.
func (h *WsHandler) subscribe(ctx context.Context, conn net.Conn) error {
	s := &subscriber{
		id:   uuid.NewString(),
		msgs: make(chan []byte, h.subscriberMessageBuffer),
		closeSlow: func() {
			err := conn.Close()
			if err != nil {
				h.Logger.Error("Unable to close socket", zap.Error(err))
			}
		},
	}
	h.addSubscriber(s)
	defer h.deleteSubscriber(s)

	h.readHandler(ctx, conn)

	go h.publish(ctx)

	for {
		select {
		case msg := <-s.msgs:
			err := writeTimeout(conn, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// publish publishes the msg to all subscribers.
// It never blocks and so messages to slow subscribers
// are dropped.
func (h *WsHandler) publish(ctx context.Context) {
	h.subscribersMu.Lock()
	defer h.subscribersMu.Unlock()

	err := h.publishLimiter.Wait(ctx)
	if err != nil {
		h.Logger.Error("Unable to publish", zap.Error(err))

		return
	}

	for s := range h.subscribers {
		select {
		case s.msgs <- h.ctrl.ToProto(s.id):
		default:
			go s.closeSlow()
		}
	}
}

// addSubscriber registers a subscriber.
func (h *WsHandler) addSubscriber(s *subscriber) {
	h.subscribersMu.Lock()
	h.subscribers[s] = struct{}{}
	h.subscribersMu.Unlock()
}

// deleteSubscriber deletes the given subscriber.
func (h *WsHandler) deleteSubscriber(s *subscriber) {
	h.subscribersMu.Lock()
	delete(h.subscribers, s)
	h.subscribersMu.Unlock()
}

func writeTimeout(conn net.Conn, msg []byte) error {
	return wsutil.WriteServerMessage(conn, ws.OpBinary, msg)
}
