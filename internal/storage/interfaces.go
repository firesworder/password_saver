package storage

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	GetUser(ctx context.Context, u User) (*User, error)
}

type RecordRepository interface {
	AddRecord(ctx context.Context, r Record, uid int) (int, error)
	UpdateRecord(ctx context.Context, r Record, uid int) error
	DeleteRecord(ctx context.Context, r Record, uid int) error
	GetAll(ctx context.Context, uid int) ([]Record, error)
}
