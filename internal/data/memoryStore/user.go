package memoryStore

import (
	"time"
	"users/internal/data"
)

// usersQ is an implementation of the data.UserQ interface
// that interacts with an in-memory data store (MemoryStore).
type usersQ struct {
	db *MemoryStore
}

// NewUserQ creates a new instance of the usersQ struct
// with a reference to the provided MemoryStore.
func NewUserQ(db *MemoryStore) data.UserQ {
	return &usersQ{
		db: db,
	}
}

// New returns a new instance of the usersQ struct.
// It is part of the data.UserQ interface.
func (u *usersQ) New() data.UserQ {
	return NewUserQ(u.db)
}

// Add adds a new user to the in-memory data store.
func (u *usersQ) Add(user data.User) (data.User, error) {
	var newUser data.User
	newUser.ID = len(u.db.users) + 1
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	u.db.users = append(u.db.users, newUser)

	return newUser, nil
}

// FindById finds a user in the in-memory data store by their ID.
func (u *usersQ) FindById(id int) (data.User, error) {
	for _, user := range u.db.users {
		if user.ID == id {
			return user, nil
		}
	}

	return data.User{}, NotFoundError
}

// GetAll retrieves a paginated list of users from the in-memory data store.
func (u *usersQ) GetAll(page int, usersPerPage int) ([]data.User, error) {
	startIndex := (page - 1) * usersPerPage
	endIndex := startIndex + usersPerPage

	if startIndex >= len(u.db.users) {
		return nil, nil
	}

	if endIndex > len(u.db.users) {
		endIndex = len(u.db.users)
	}

	return u.db.users[startIndex:endIndex], nil
}

// Update updates the information of a user in the in-memory data store.
func (u *usersQ) Update(id int, userUpdate data.User) (data.User, error) {
	for i, user := range u.db.users {
		if user.ID == id {
			if userUpdate.Name != "" {
				u.db.users[i].Name = userUpdate.Name
			}
			if userUpdate.Email != "" {
				u.db.users[i].Email = userUpdate.Email
			}
			u.db.users[i].UpdatedAt = time.Now()

			return u.db.users[i], nil
		}
	}

	return data.User{}, NotFoundError
}

// Delete removes a user from the in-memory data store by their ID.
func (u *usersQ) Delete(id int) error {
	for i, user := range u.db.users {
		if user.ID == id {
			// Deleting user from the slice using slice indexes and spread.
			u.db.users = append(u.db.users[:i], u.db.users[i+1:]...)
			return nil
		}
	}
	return NotFoundError
}
