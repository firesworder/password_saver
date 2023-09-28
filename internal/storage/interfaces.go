package storage

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	GetUser(ctx context.Context, u User) (*User, error)
}

type TextDataRepository interface {
	AddTextData(ctx context.Context, td TextData, u *User) (int, error)
	UpdateTextData(ctx context.Context, td TextData, u *User) error
	DeleteTextData(ctx context.Context, td TextData, u *User) error
	GetAllRecords(ctx context.Context, u *User) ([]TextData, error)
}

type BankDataRepository interface {
	AddBankData(ctx context.Context, bd BankData, u *User) (int, error)
	UpdateBankData(ctx context.Context, bd BankData, u *User) error
	DeleteBankData(ctx context.Context, bd BankData, u *User) error
	GetAllRecords(ctx context.Context, u *User) ([]BankData, error)
}

type BinaryDataRepository interface {
	AddBinaryData(ctx context.Context, bd BinaryData, u *User) (int, error)
	UpdateBinaryData(ctx context.Context, bd BinaryData, u *User) error
	DeleteBinaryData(ctx context.Context, bd BinaryData, u *User) error
	GetAllRecords(ctx context.Context, u *User) ([]BinaryData, error)
}

type RecordRepository interface {
	AddRecord(ctx context.Context, r Record, uid int) (int, error)
	UpdateRecord(ctx context.Context, r Record, uid int) error
	DeleteRecord(ctx context.Context, r Record, uid int) error
	GetAll(ctx context.Context, uid int) ([]Record, error)
}
