package models

import "database/sql"

type FolderStore interface {
	GetAllFolder() ([]*Folder, error)
	GetFolderByID(id int) (*Folder, error)
	GetFolderByName(name string) (*Folder, error)
	CreateFolder(Folder *CreateFolderPayload) (int64, error)
	UpdateFolder(Folder *UpdateFolderPayload) error
	DeleteFolder(id int) error
}

type Folder struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateFolderPayload struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=30"`
}

type UpdateFolderPayload struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=30"`
}
