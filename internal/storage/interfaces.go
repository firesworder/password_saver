package storage

import "context"

type User struct {
	Login, HashedPassword string
}

type UserRepository interface {
	CreateUser(ctx context.Context, u User) error
	GetUser(ctx context.Context, u User) (*User, error)
}
