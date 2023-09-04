package users

import (
	"context"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
)

var ErrUserExist = errors.New("user already exist")
var ErrUserNotExist = errors.New("user not exist")

type MockUser struct {
	Users      map[string]storage.User
	LastUsedID int
}

func (m *MockUser) CreateUser(ctx context.Context, u storage.User) (*storage.User, error) {
	if _, ok := m.Users[u.Login]; ok {
		return nil, ErrUserExist
	}
	m.LastUsedID++
	u.ID = m.LastUsedID
	m.Users[u.Login] = u
	return &u, nil
}

func (m *MockUser) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	user, ok := m.Users[u.Login]
	if !ok {
		return nil, ErrUserNotExist
	}
	return &user, nil
}
