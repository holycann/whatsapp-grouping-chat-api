package models

import (
	"database/sql"
)

type ChatStore interface {
	GetAllChat() ([]*Chat, error)
	GetChatByID(id int) (*Chat, error)
	CreateChat(Chat *Chat) error
	UpdateChat(Chat *Chat) error
	DeleteChat(id int) error
}

type Chat struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	Message   string       `json:"message"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateChatPayload struct {
	UserID  int    `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type UpdateChatPayload struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
