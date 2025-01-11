package models

import "google.golang.org/protobuf/types/known/timestamppb"

type ChatStore interface {
	GetAllChat() ([]*Chat, error)
	GetChatByID(id int) (*Chat, error)
	CreateChat(Chat *Chat) error
	UpdateChat(Chat *Chat) error
	DeleteChat(id int) error
}

type Chat struct {
	ID        int                   `json:"id"`
	UserID    int                   `json:"user_id"`
	Message   string                `json:"message"`
	CreatedAt timestamppb.Timestamp `json:"created_at"`
	UpdatedAt timestamppb.Timestamp `json:"updated_at"`
}

type CreateChatPayload struct {
	Message   string                `json:"name" validate:"required"`
	CreatedAt timestamppb.Timestamp `json:"created_at"`
}

type UpdateChatPayload struct {
	ID        int                   `json:"id"`
	Message   string                `json:"name" validate:"required"`
	UpdatedAt timestamppb.Timestamp `json:"updated_at"`
}
