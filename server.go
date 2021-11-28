package dcc

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var wsServer = NewWebsocketServer()
var ctrl *Controller

func Serve(c *Controller) {
	log.Println("Listening on :3000...")
	ctrl = c
	go wsServer.Run()
	go sendState(c)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func sendState(ctrl *Controller) {
	for {
		time.Sleep(3 * time.Second)
		wsServer.broadcastToClients(ctrl.ToJson())
	}
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer() *WsServer {
	return &WsServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
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
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
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
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(*http.Request) bool { return true },
}

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
}

func newClient(conn *websocket.Conn, wsServer *WsServer) *Client {
	return &Client{
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256),
	}

}

func (client *Client) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		var cj *ControllerJson

		json.Unmarshal(jsonMessage, &cj)

		for _, loco := range ctrl.Locos() {
			loco.Speed = cj.Locomotives[loco.Name].Speed
			loco.Direction = cj.Locomotives[loco.Name].Direction
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

			loco.Apply()
		}

		client.wsServer.broadcast <- jsonMessage
	}

}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	close(client.send)
	client.conn.Close()
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer)

	go client.writePump()
	go client.readPump()

	wsServer.register <- client
}
