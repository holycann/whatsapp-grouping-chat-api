package models

type ChatStore interface {
	GetAllChat() ([]*Chat, error)
	GetChatByID(id int) (*Chat, error)
	CreateChat(Chat *Chat) error
	UpdateChat(Chat *Chat) error
	DeleteChat(id int) error
}

type Chat struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type CreateChatPayload struct {
	Message string `json:"name" validate:"required"`
}

type UpdateChatPayload struct {
	Message string `json:"name" validate:"required"`
}
