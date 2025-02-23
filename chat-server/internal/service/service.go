package service

import (
	"time"

	"github.com/yourusername/chat/internal/model"
	"github.com/yourusername/chat/internal/repository"
	"github.com/yourusername/chat/internal/websocket"
)

type ChatService struct {
	repo *repository.PostgresRepository
	hub  *websocket.Hub
}

func NewChatService(repo *repository.PostgresRepository, hub *websocket.Hub) *ChatService {
	return &ChatService{
		repo: repo,
		hub:  hub,
	}
}

func (s *ChatService) SaveMessage(username, content string) (*model.Message, error) {
	msg := &model.Message{
		Username:  username,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveMessage(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *ChatService) GetMessages(limit int) ([]model.Message, error) {
	return s.repo.GetMessages(limit)
}

func (s *ChatService) GetHub() *websocket.Hub {
	return s.hub
}
