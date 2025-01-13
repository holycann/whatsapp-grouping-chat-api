package models

type UserStore interface {
	GetAllUser() ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *UpdateUserPayload) error
	DeleteUser(id int) error
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	ImageURL    string `json:"image"`
}

type CreateUserPayload struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	ImageURL    string `json:"image_url"`
}

type UpdateUserPayload struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	ImageURL    string `json:"image_url"`
}
