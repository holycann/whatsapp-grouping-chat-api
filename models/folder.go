package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FolderStore interface {
	GetFolderByName(name string) (*Folder, error)
	GetFolderByID(id int) (*Folder, error)
	CreateFolder(Folder *CreateFolderPayload) error
	UpdateFolder(Folder *UpdateFolderPayload) error
	DeleteFolder(id int) error
}

type Folder struct {
	ID        int                   `json:"id"`
	ChatID    int                   `json:"chat_id"`
	Name      string                `json:"name"`
	CreatedAt timestamppb.Timestamp `json:"created_at"`
	UpdatedAt timestamppb.Timestamp `json:"updated_at"`
}

type CreateFolderPayload struct {
	ChatID    int                   `json:"chat_id" validate:"required"`
	Name      string                `json:"name" validate:"required, min=3, max=30"`
	CreatedAt timestamppb.Timestamp `json:"created_at" validate:"required"`
}

type UpdateFolderPayload struct {
	ID        int                   `json:"id" validate:"required"`
	ChatID    int                   `json:"chat_id" validate:"required"`
	Name      string                `json:"name" validate:"required, min=3, max=30"`
	UpdatedAt timestamppb.Timestamp `json:"updated_at" validate:"required"`
}
