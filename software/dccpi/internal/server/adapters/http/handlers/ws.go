package handlers

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/alexbegoon/go-dcc/internal/controller"
	controllerProto "github.com/alexbegoon/go-dcc/internal/pb/build/go/controller"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"

	"nhooyr.io/websocket"
)

const (
	WriteTimeout       = 5 * time.Second
	MaxMessagesInQueue = 16
	MaxMessageLength   = 8192
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

func (h *WsHandler) readHandler(ctx context.Context, c *websocket.Conn) {
	go func() {
		for {
			typ, r, err := c.Read(ctx)
			if err != nil {
				h.Logger.Error("Error while reading from socket", zap.Error(err))
			}

			if typ != websocket.MessageBinary {
				h.Logger.Error("Wrong data type received", zap.Any("type", typ))
			}

			var cProto controllerProto.Controller

			err = proto.Unmarshal(r, &cProto)
			if err != nil {
				h.Logger.Error("Unable to Unmarshal", zap.Error(err))

				return
			}

			err = h.ctrl.Handle(&cProto)
			if err != nil {
				h.Logger.Error("Unable to Handle command from subscriber", zap.Error(err))

				return
			}

			h.publish(ctx)
		}
	}()
}

// subscribeHandler accepts the WebSocket connection and then subscribes
// it to all future messages.
func (h *WsHandler) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		h.Logger.Error("Unable to accept handshake", zap.Error(err))

		return
	}
	defer func(c *websocket.Conn, code websocket.StatusCode, reason string) {
		err = c.Close(code, reason)
		if err != nil {
			h.Logger.Error("Unable to close connection", zap.Error(err))
		}
	}(c, websocket.StatusInternalError, "")

	err = h.subscribe(r.Context(), c)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		h.Logger.Error("Unable to subscribe", zap.Error(err))

		return
	}

	h.readHandler(r.Context(), c)

	h.publish(r.Context())
}

// subscribe subscribes the given WebSocket to all broadcast messages.
// It creates a subscriber with a buffered msgs chan to give some room to slower
// connections and then registers the subscriber. It then listens for all messages
// and writes them to the WebSocket. If the context is canceled or
// an error occurs, it returns and deletes the subscription.
//
// It uses CloseRead to keep reading from the connection to process control
// messages and cancel the context if the connection drops.
func (h *WsHandler) subscribe(ctx context.Context, c *websocket.Conn) error {
	ctx = c.CloseRead(ctx)

	s := &subscriber{
		id:   uuid.NewString(),
		msgs: make(chan []byte, h.subscriberMessageBuffer),
		closeSlow: func() {
			err := c.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
			if err != nil {
				h.Logger.Error("Unable to close socket", zap.Error(err))
			}
		},
	}
	h.addSubscriber(s)
	defer h.deleteSubscriber(s)

	for {
		select {
		case msg := <-s.msgs:
			err := writeTimeout(ctx, WriteTimeout, c, msg)
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

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageBinary, msg)
}
