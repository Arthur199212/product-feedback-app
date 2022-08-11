package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for test purposes
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	l    *logrus.Logger
	conn *websocket.Conn
	// Handles clients that was disconnected
	unregisterCallback func(client *Client)
	// Buffered channel of outbound messages
	Send chan []byte
}

func NewClient(
	l *logrus.Logger,
	w http.ResponseWriter,
	r *http.Request,
	unregisterCallback func(client *Client),
) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		l:                  l,
		conn:               conn,
		unregisterCallback: unregisterCallback,
		Send:               make(chan []byte, 256),
	}

	return client, err
}

func (c *Client) ReadPump() {
	defer func() {
		c.unregisterCallback(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			isUnexpectedErr := websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			)
			if isUnexpectedErr {
				log.Printf("error: %v", err)
			}
			break
		}
		// We don't expect messages from client as of now
		// just log them for test purposes
		c.l.Println("message from client:", msg)
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Channel was closed
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.l.Error(err)
				return
			}

			if _, err := w.Write(msg); err != nil {
				c.l.Error(err)
				return
			}

			if err := w.Close(); err != nil {
				c.l.Error(err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.l.Error(err)
				return
			}
		}
	}
}
