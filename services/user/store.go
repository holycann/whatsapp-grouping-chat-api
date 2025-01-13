package user

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

func scanRowIntoUser(row *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetAllUser() (users []*models.User, err error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (s *Store) GetUserByUsername(name string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Name == "" {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, phone_number, image_url) VALUES ($1, $2, $3)", user.Name, user.PhoneNumber, user.ImageURL)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(user *models.UpdateUserPayload) error {
	_, err := s.db.Exec("UPDATE users SET name = $1, phone_number = $2, image_url = $3 WHERE id = $4", user.Name, user.PhoneNumber, user.ImageURL, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
