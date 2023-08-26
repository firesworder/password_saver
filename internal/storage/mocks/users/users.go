package users

import (
	"context"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
)

var ErrUserExist = errors.New("user already exist")
var ErrUserNotExist = errors.New("user not exist")

type MockUser struct {
	Users map[string]storage.User
}

func (m *MockUser) CreateUser(ctx context.Context, u storage.User) error {
	if _, ok := m.Users[u.Login]; ok {
		return ErrUserExist
	}
	m.Users[u.Login] = u
	return nil
}

func (m *MockUser) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	user, ok := m.Users[u.Login]
	if !ok {
		return nil, ErrUserNotExist
	}
	return &user, nil
}
