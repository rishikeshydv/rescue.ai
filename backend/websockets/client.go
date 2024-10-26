package websockets

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writingWaitTime = 10 * time.Second
	pongTime        = 60 * time.Second
	pingTime        = (pongTime * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *Hub
}

// reading messages from client
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadDeadline(time.Now().Add(pongTime))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongTime))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.broadcast <- msg
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingTime)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {

	}
}
func ServerWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	newClient := &Client{
		conn: conn,
		send: make(chan []byte),
		hub:  hub,
	}
	newClient.hub.register <- newClient

	go newClient.WritePump()
	go newClient.ReadPump()
}
