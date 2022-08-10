package notifier

import (
	"encoding/json"
	"product-feedback/ws"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type NotifierService interface {
	BroadcastMessage(et EventType, sub SubjectType, id int)
	Register(client *ws.Client)
}

// notifierService maintains the set of active clients
// and broadcasts messages to the clients
type notifierService struct {
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

func NewNotifierSerivice() NotifierService {
	hub := &notifierService{
		broadcast:  make(chan []byte),
		register:   make(chan *ws.Client),
		unregister: make(chan *ws.Client),
		clients:    make(map[*ws.Client]bool),
	}
	go hub.run()
	return hub
}

type EventType string
type SubjectType string

const (
	CreateEvent EventType = "create"
	UpdateEvent EventType = "update"
	DeleteEvent EventType = "delete"

	SubjectFeedback SubjectType = "feedback"
	SubjectComment  SubjectType = "comment"
	SubjectVote     SubjectType = "vote"
)

type Message struct {
	// enum: create, update, delete
	EventType EventType `json:"eventType"`
	// enum: feedback, comment, vote
	Subject SubjectType `json:"subject"`
	// id of the subject
	Id int `json:"id"`
}

func (s *notifierService) BroadcastMessage(
	et EventType,
	sub SubjectType,
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

func (s *notifierService) Register(client *ws.Client) {
	s.register <- client
}

func (s *notifierService) run() {
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
