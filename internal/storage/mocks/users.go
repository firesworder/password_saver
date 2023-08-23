package mocks

import (
	"context"
	"errors"
	"password_saver/internal/storage"
)

var ErrUserExist = errors.New("user already exist")
var ErrUserNotExist = errors.New("user not exist")

type MockUser struct {
	users map[string]storage.User
}

func (m *MockUser) CreateUser(ctx context.Context, u storage.User) error {
	if _, ok := m.users[u.Login]; ok {
		return ErrUserExist
	}
	m.users[u.Login] = u
	return nil
}

func (m *MockUser) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	user, ok := m.users[u.Login]
	if !ok {
		return nil, ErrUserNotExist
	}
	return &user, nil
}
