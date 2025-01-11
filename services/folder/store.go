package folder

import (
	"database/sql"
	"fmt"

	"github.com/holycann/whatsapp-grouping-chat-api/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowIntoFolder(row *sql.Rows) (*models.Folder, error) {
	folder := new(models.Folder)

	err := row.Scan(
		&folder.ID,
		&folder.ChatID,
		&folder.Name,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (s *Store) GetFolderByName(name string) (*models.Folder, error) {
	rows, err := s.db.Query("SELECT * FROM folders WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	u := new(models.Folder)
	for rows.Next() {
		u, err = scanRowIntoFolder(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Name == "" {
		return nil, fmt.Errorf("Folder not found")
	}

	return u, nil
}

func (s *Store) GetFolderByID(id int) (*models.Folder, error) {
	rows, err := s.db.Query("SELECT * FROM folders WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(models.Folder)
	for rows.Next() {
		u, err = scanRowIntoFolder(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("Folder not found")
	}

	return u, nil
}

func (s *Store) CreateFolder(folder *models.Folder) error {
	_, err := s.db.Exec("INSERT INTO folders (chat_id, name, created_at VALUES (?, ?, ?)", folder.ChatID, folder.Name, folder.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateFolder(folder *models.UpdateFolderPayload) error {
	_, err := s.db.Exec("UPDATE folders SET `chat_id` = ?, `name` = ?, `updated_at` = ? WHERE id = ?", folder.ChatID, folder.Name, folder.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteFolder(id int) error {
	_, err := s.db.Exec("DELETE FROM folders WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
