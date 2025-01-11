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
		&chat.Message,
	)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Store) GetAllChat() ([]*models.Chat, error) {
	rows, err := s.db.Query("SELECT * FROM chat")
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
	rows, err := s.db.Query("SELECT * FROM chat WHERE id = ?", id)
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
	_, err := s.db.Exec("INSERT INTO chat (`message`) VALUES (?)", chat.Message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateChat(chat *models.Chat) error {
	_, err := s.db.Exec("UPDATE chat SET `message` = ? WHERE id = ?", chat.Message, chat.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteChat(id int) error {
	_, err := s.db.Exec("DELETE FROM chat WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
