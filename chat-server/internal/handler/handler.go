package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yourusername/chat/internal/service"
	ws "github.com/yourusername/chat/internal/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // For development purposes
	},
}

type Handler struct {
	service *service.ChatService
}

func NewHandler(service *service.ChatService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username := c.Query("username")
	if username == "" {
		username = "anonymous"
	}

	client := &ws.Client{
		Hub:      h.service.GetHub(),
		Conn:     conn,
		Send:     make(chan []byte, 256),
		Username: username,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

func (h *Handler) GetMessages(c *gin.Context) {
	messages, err := h.service.GetMessages(100) // Get last 100 messages
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
