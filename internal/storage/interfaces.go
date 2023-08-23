package storage

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, u User) error
	GetUser(ctx context.Context, u User) (*User, error)
}

type TextDataRepository interface {
	AddTextData(ctx context.Context, td TextData) error
	UpdateTextData(ctx context.Context, td TextData) error
	DeleteTextData(ctx context.Context, td TextData) error
}
