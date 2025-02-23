package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yourusername/chat/config"
	"github.com/yourusername/chat/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(cfg config.DatabaseConfig) (*PostgresRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) SaveMessage(msg *model.Message) error {
	query := `
        INSERT INTO messages (username, content, created_at)
        VALUES ($1, $2, $3)
        RETURNING id`

	return r.db.QueryRow(query, msg.Username, msg.Content, msg.CreatedAt).Scan(&msg.ID)
}

func (r *PostgresRepository) GetMessages(limit int) ([]model.Message, error) {
	query := `
        SELECT id, username, content, created_at
        FROM messages
        ORDER BY created_at DESC
        LIMIT $1`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.ID, &msg.Username, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
