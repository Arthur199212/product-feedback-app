package notifier

import (
	"product-feedback/ws"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NotifierHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type notifierHandler struct {
	l       *logrus.Logger
	service NotifierService
}

func NewNotifierHandler(
	l *logrus.Logger,
	service NotifierService,
) NotifierHandler {
	return &notifierHandler{
		l:       l,
		service: service,
	}
}

func (h *notifierHandler) AddRoutes(r *gin.RouterGroup) {
	r.GET("/ws", h.serveWS)
}

func (h *notifierHandler) serveWS(c *gin.Context) {
	client, err := ws.NewClient(h.l, c.Writer, c.Request)
	if err != nil {
		h.l.Error(err)
		return
	}

	h.service.Register(client)

	go client.WritePump()
}
