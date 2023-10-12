package storage

import "context"

// UserRepository интерфейс доступа к данным пользователей.
type UserRepository interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	GetUser(ctx context.Context, u User) (*User, error)
}

// RecordRepository интерфей доступа к данным записей пользователей.
type RecordRepository interface {
	AddRecord(ctx context.Context, r Record, uid int) (int, error)
	UpdateRecord(ctx context.Context, r Record, uid int) error
	DeleteRecord(ctx context.Context, r Record, uid int) error
	GetAll(ctx context.Context, uid int) ([]Record, error)
}
