package websockets

import (
	"backend/models"
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
func (c *Client) ClientMessage() {
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
		AddPoliceMessages(
			&models.Message{
				MessageDepartment: "Police",
				MessageType:       "Sent",
				DeliveredTime:     time.Now(),
				MessageContent:    string(msg),
				SenderUser:        c.conn.RemoteAddr().String()})

	}
}

// when one user sends a message, it goes to the server only
// so its the server's job to send that message to all the other users
// thats what 'WritePump' do
func (c *Client) ServerMessage() {
	ticker := time.NewTicker(pingTime)
	defer func() {
		c.conn.Close()
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.SetWriteDeadline(time.Now().Add(writingWaitTime))
			writer, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			writer.Write(message)
			AddPoliceMessages(
				&models.Message{
					MessageDepartment: "Police",
					MessageType:       "Received",
					DeliveredTime:     time.Now(),
					MessageContent:    string(message),
					SenderUser:        c.conn.RemoteAddr().String()})

			//handling messages in the queue
			totalMessages := len(c.send)
			for i := 0; i < totalMessages; i++ {
				writer.Write([]byte{'\n'}) //this is for a new line
				writer.Write(<-c.send)
			}

			//sending a ping at every ping interval
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writingWaitTime))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
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
	json.NewEncoder(w).Encode(map[string]string{"Status": "Client Successfully Assigned a Websocket"})
	go newClient.ClientMessage()
	go newClient.ServerMessage()
}
