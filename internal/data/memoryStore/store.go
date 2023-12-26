// Package memoryStore provides an in-memory database imitation for user data.
package memoryStore

import (
	"errors"
	"users/internal/data"
)

var (
	NotFoundError = errors.New("not found")
)

type MemoryStore struct {
	users []data.User
}

func NewMemoryStore() *MemoryStore {
	var users []data.User
	return &MemoryStore{
		users,
	}
}
