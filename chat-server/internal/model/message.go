package model

import "time"

type Message struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	Username string `json:"username"`
}
