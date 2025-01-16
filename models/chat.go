package models

import (
	"database/sql"

)

type ChatStore interface {
	GetAllChat() ([]*Chat, error)
	GetChatByID(id int) (*Chat, error)
	CreateChat(CreateChatPayload *CreateChatPayload) error
	UpdateChat(UpdateChatPayload *UpdateChatPayload) error
	DeleteChat(id int) error
}

type Chat struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	FolderID  sql.NullInt16 `json:"folder_id"`
	Message   string        `json:"message"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type CreateChatPayload struct {
	UserID   int    `json:"user_id" validate:"required"`
	FolderID int    `json:"folder_id"`
	Message  string `json:"message" validate:"required"`
}

type UpdateChatPayload struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id" validate:"required"`
	FolderID int    `json:"folder_id"`
	Message  string `json:"message" validate:"required"`
}
