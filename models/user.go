package models

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *UpdateUserPayload) error
	DeleteUser(id int) error
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	ImageURL    string `json:"image"`
}

type CreateUserPayload struct {
	Username    string `json:"username" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	ImageURL    string `json:"image"`
}

type UpdateUserPayload struct {
	ID          int    `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required, min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required, min=3,max=20"`
	ImageURL    string `json:"image"`
}
