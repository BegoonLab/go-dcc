package server

import (
	"encoding/json"
	"log"
	"net/http"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/alexbegoon/go-dcc/software/dccpi/internal/controller"
	"github.com/gorilla/websocket"
)

var (
	wsServer *WsServer
	ctrl     *controller.Controller
)

func Serve(c *controller.Controller) *http.Server {
	wsServer = NewWebsocketServer(c.Log)
	log.Println("Listening on :3000...")
	ctrl = c
	ctrl.Stop()
	go wsServer.Run()
	go sendState(c)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("../../frontend/public"))
	http.Handle("/", fs)

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: 3 * time.Second,  // nolint:gomnd
		ReadTimeout:       10 * time.Second, // nolint:gomnd
		WriteTimeout:      10 * time.Second, // nolint:gomnd
	}

	if err := server.ListenAndServe(); err != nil {
		c.Log.Error("Unable to start server", zap.Error(err))
	}

	return server
}

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	log        *zap.Logger
}

func sendState(ctrl *controller.Controller) {
	for {
		time.Sleep(3 * time.Second) // nolint:forbidigo,gomnd
		wsServer.broadcastToClients(ctrl.ToJSON())
	}
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer(log *zap.Logger) *WsServer {
	return &WsServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		log:        log,
	}
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {
	for {
		select {

		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)

		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *WsServer) registerClient(client *Client) {
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *Client) {
	delete(server.clients, client)
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10 // nolint:gomnd

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var newline = []byte{'\n'}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096, // nolint:gomnd
	WriteBufferSize: 4096, // nolint:gomnd
	CheckOrigin:     func(*http.Request) bool { return true },
}

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
	log      *zap.Logger
}

func newClient(conn *websocket.Conn, wsServer *WsServer, log *zap.Logger) *Client {
	return &Client{
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256), // nolint:gomnd
		log:      log,
	}
}

func (client *Client) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	err := client.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		client.log.Error("Unable to set read deadline", zap.Error(err))
	}

	client.conn.SetPongHandler(func(string) error { return client.conn.SetReadDeadline(time.Now().Add(pongWait)) })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			client.log.Error("Unable to Read Message", zap.Error(err))
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				client.log.Error("unexpected close error", zap.Error(err))
			}

			break
		}

		var cj *controller.ControllerJSON

		err = json.Unmarshal(jsonMessage, &cj)
		if err != nil {
			client.log.Error("Unable to Unmarshal", zap.Error(err))
		}

		if !cj.Started && ctrl.IsStarted() {
			ctrl.Stop()
		}

		if cj.Started && !ctrl.IsStarted() {
			ctrl.Start()
		}

		if cj.Reboot {
			log.Println("Rebooting the system...")

			syscall.Sync()
			err = syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
			if err != nil {
				log.Printf("Reboot failed: %v", err)
			}
		}

		if cj.Poweroff {
			client.log.Info("Shutdown the system...", zap.Error(err))

			syscall.Sync()
			err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
			if err != nil {
				client.log.Error("Shutdown failed", zap.Error(err))
			}
		}

		for _, loco := range ctrl.Locos() {
			loco.Speed = cj.Locomotives[loco.Name].Speed
			loco.Direction = cj.Locomotives[loco.Name].Direction
			loco.Enabled = cj.Locomotives[loco.Name].Enabled
			loco.Fl = cj.Locomotives[loco.Name].Fl
			loco.F1 = cj.Locomotives[loco.Name].F1
			loco.F2 = cj.Locomotives[loco.Name].F2
			loco.F3 = cj.Locomotives[loco.Name].F3
			loco.F4 = cj.Locomotives[loco.Name].F4
			loco.F5 = cj.Locomotives[loco.Name].F5
			loco.F6 = cj.Locomotives[loco.Name].F6
			loco.F7 = cj.Locomotives[loco.Name].F7
			loco.F8 = cj.Locomotives[loco.Name].F8
			loco.F9 = cj.Locomotives[loco.Name].F9
			loco.F10 = cj.Locomotives[loco.Name].F10
			loco.F11 = cj.Locomotives[loco.Name].F11
			loco.F12 = cj.Locomotives[loco.Name].F12
			loco.F13 = cj.Locomotives[loco.Name].F13
			loco.F14 = cj.Locomotives[loco.Name].F14
			loco.F15 = cj.Locomotives[loco.Name].F15
			loco.F16 = cj.Locomotives[loco.Name].F16
			loco.F17 = cj.Locomotives[loco.Name].F17
			loco.F18 = cj.Locomotives[loco.Name].F18
			loco.F19 = cj.Locomotives[loco.Name].F19
			loco.F20 = cj.Locomotives[loco.Name].F20
			loco.F21 = cj.Locomotives[loco.Name].F21
			loco.F22 = cj.Locomotives[loco.Name].F22
			loco.F23 = cj.Locomotives[loco.Name].F23
			loco.F24 = cj.Locomotives[loco.Name].F24
			loco.F25 = cj.Locomotives[loco.Name].F25
			loco.F26 = cj.Locomotives[loco.Name].F26
			loco.F27 = cj.Locomotives[loco.Name].F27
			loco.F28 = cj.Locomotives[loco.Name].F28

			loco.Apply()
		}

		client.wsServer.broadcast <- jsonMessage
	}
}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := client.conn.Close()
		if err != nil {
			return
		}
	}()
	for {
		select {
		case message, ok := <-client.send:
			err := client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				client.log.Error("Unable to set write deadline", zap.Error(err))
			}
			if !ok {
				// The WsServer closed the channel.
				err = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					client.log.Error("Unable to write message", zap.Error(err))
				}

				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				client.log.Error("Unable to write", zap.Error(err))
			}

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				_, err = w.Write(newline)
				if err != nil {
					client.log.Error("Unable to write", zap.Error(err))
				}
				if err != nil {
					return
				}
				_, err = w.Write(<-client.send)
				if err != nil {
					client.log.Error("Unable to write", zap.Error(err))
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			err := client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				client.log.Error("Unable to set write deadline", zap.Error(err))
			}
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	close(client.send)
	if err := client.conn.Close(); err != nil {
		client.log.Error("Unable to close connection", zap.Error(err))
	}
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

		return
	}

	client := newClient(conn, wsServer, wsServer.log)

	go client.writePump()
	go client.readPump()

	wsServer.register <- client
}
