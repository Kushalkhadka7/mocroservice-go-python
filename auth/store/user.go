package store

import (
	"auth/model"
)

// UserStore interface to statisfy.
type UserStore interface {
	Save(user *model.User) (*model.User, error)
}

// CreateUserStore initializes new user store.
type CreateUserStore struct{}

// NewUesrStore creates new user store.
func NewUesrStore() *CreateUserStore {
	return &CreateUserStore{}
}

// Save user to db.
func (store *CreateUserStore) Save(user *model.User) (*model.User, error) {
	return user, nil
}
