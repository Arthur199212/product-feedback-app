package notifier

import (
	"encoding/json"
	"product-feedback/ws"

	"github.com/sirupsen/logrus"
)

// NotifierService maintains the set of active clients
// and broadcasts messages to the clients
type NotifierService struct {
	l *logrus.Logger
	// Registered clients
	clients map[*ws.Client]bool
	// List of messages to broadcast
	broadcast chan []byte
	// Register clients
	register chan *ws.Client
	// Unregister clients
	unregister chan *ws.Client
}

func NewNotifierSerivice() *NotifierService {
	hub := &NotifierService{
		broadcast:  make(chan []byte),
		register:   make(chan *ws.Client),
		unregister: make(chan *ws.Client),
		clients:    make(map[*ws.Client]bool),
	}
	go hub.run()
	return hub
}

type eventType string
type subjectType string

const (
	CreateEvent eventType = "create"
	UpdateEvent eventType = "update"
	DeleteEvent eventType = "delete"

	SubjectFeedback subjectType = "feedback"
	SubjectComment  subjectType = "comment"
	SubjectVote     subjectType = "vote"
)

type Message struct {
	// enum: create, update, delete
	EventType eventType `json:"eventType"`
	// enum: feedback, comment, vote
	Subject subjectType `json:"subject"`
	// id of the subject
	Id int `json:"id"`
}

func (s *NotifierService) BroadcastMessage(
	et eventType,
	sub subjectType,
	id int,
) {
	msg := Message{
		EventType: et,
		Id:        id,
		Subject:   sub,
	}
	msgByte, err := json.Marshal(msg)
	if err != nil {
		s.l.Warn(err)
		return
	}
	s.broadcast <- msgByte
}

func (s *NotifierService) run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.Send)
			}
		case msg := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(s.clients, client)
				}
			}
		}
	}
}
