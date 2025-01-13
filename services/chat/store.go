package chat

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

func scanRowIntoChat(row *sql.Rows) (*models.Chat, error) {
	chat := new(models.Chat)

	err := row.Scan(
		&chat.ID,
		&chat.UserID,
		&chat.Message,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Store) GetAllChat() ([]*models.Chat, error) {
	rows, err := s.db.Query("SELECT * FROM chats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*models.Chat
	for rows.Next() {
		chat, err := scanRowIntoChat(rows)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

func (s *Store) GetChatByID(id int) (*models.Chat, error) {
	rows, err := s.db.Query("SELECT * FROM chats WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	r := new(models.Chat)
	for rows.Next() {
		r, err = scanRowIntoChat(rows)
		if err != nil {
			return nil, err
		}
	}

	if r.ID == 0 {
		return nil, fmt.Errorf("Chat not found")
	}

	return r, nil
}

func (s *Store) CreateChat(chat *models.Chat) error {
	_, err := s.db.Exec("INSERT INTO chats (user_id, message) VALUES ($1, $2)", chat.UserID, chat.Message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateChat(chat *models.Chat) error {
	_, err := s.db.Exec("UPDATE chats SET user_id = $1, message = $2 WHERE id = $3", chat.UserID, chat.Message, chat.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteChat(id int) error {
	_, err := s.db.Exec("DELETE FROM chats WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
