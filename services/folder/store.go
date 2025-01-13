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
		&folder.Name,
		&folder.ChatID,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (s *Store) GetAllFolder() (folder []*models.Folder, err error) {
	rows, err := s.db.Query("SELECT * FROM folders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f, err := scanRowIntoFolder(rows)
		if err != nil {
			return nil, err
		}
		folder = append(folder, f)
	}

	return folder, nil
}

func (s *Store) GetFolderByName(name string) (*models.Folder, error) {
	rows, err := s.db.Query("SELECT * FROM folders WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	f := new(models.Folder)
	for rows.Next() {
		f, err = scanRowIntoFolder(rows)
		if err != nil {
			return nil, err
		}
	}

	if f.Name == "" {
		return nil, fmt.Errorf("Folder not found")
	}

	return f, nil
}

func (s *Store) GetFolderByID(id int) (*models.Folder, error) {
	rows, err := s.db.Query("SELECT * FROM folders WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	f := new(models.Folder)
	for rows.Next() {
		f, err = scanRowIntoFolder(rows)
		if err != nil {
			return nil, err
		}
	}

	if f.ID == 0 {
		return nil, fmt.Errorf("Folder not found")
	}

	return f, nil
}

func (s *Store) CreateFolder(folder *models.CreateFolderPayload) error {
	_, err := s.db.Exec("INSERT INTO folders (chat_id, name) VALUES ($1, $2)", folder.ChatID, folder.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateFolder(folder *models.UpdateFolderPayload) error {
	_, err := s.db.Exec("UPDATE folders SET chat_id = $1, name = $2 WHERE id = $3", folder.ChatID, folder.Name, folder.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteFolder(id int) error {
	_, err := s.db.Exec("DELETE FROM folders WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
