package models

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *UpdateUserPayload) error
	DeleteUser(id int) error
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

type CreateUserPayload struct {
	Fullname string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,number,min=10,max=20"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	RoleID   int    `json:"role_id" validate:"required"`
}

type UpdateUserPayload struct {
	Fullname string `json:"fullname"`
	Username string `json:"username" validate:"min=3,max=30"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone" validate:"number,min=10,max=20"`
	Password string `json:"password" validate:"min=8,max=32"`
	RoleID   int    `json:"role_id"`
}
