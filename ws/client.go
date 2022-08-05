package ws

import (
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
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// allow all origins for test purposes
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	l    *logrus.Logger
	conn *websocket.Conn
	// Buffered channel of outbound messages
	Send chan []byte
}

func NewClient(
	l *logrus.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		l:    l,
		conn: conn,
		Send: make(chan []byte, 256),
	}

	return client, err
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
