package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/chat/config"
	"github.com/yourusername/chat/internal/handler"
	"github.com/yourusername/chat/internal/repository"
	"github.com/yourusername/chat/internal/service"
	"github.com/yourusername/chat/internal/websocket"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize repository
	repo, err := repository.NewPostgresRepository(cfg.DB)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	// Initialize WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// Initialize service
	svc := service.NewChatService(repo, hub)

	// Initialize handler
	h := handler.NewHandler(svc)

	// Setup router
	r := gin.Default()

	// Serve static files for chat UI
	r.Static("/static", "./public")

	// WebSocket endpoint
	r.GET("/api/ws", h.HandleWebSocket)

	// REST endpoints for message history
	r.GET("/messages", h.GetMessages)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
