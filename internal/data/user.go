package data

import "time"

// UserQ is an interface that represents operations on user data.
type UserQ interface {
	New() UserQ
	Add(user User) (User, error)
	FindById(id int) (User, error)
	GetAll(page int, usersPerPage int) ([]User, error)
	Update(id int, userUpdate User) (User, error)
	Delete(id int) error
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
