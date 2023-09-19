package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
)

type IServer interface {
	RegisterUser(ctx context.Context, user storage.User) (string, error)
	LoginUser(ctx context.Context, user storage.User) (string, error)

	AddTextData(ctx context.Context, textData storage.TextData) (int, error)
	UpdateTextData(ctx context.Context, textData storage.TextData) error
	DeleteTextData(ctx context.Context, textData storage.TextData) error

	AddBankData(ctx context.Context, bankData storage.BankData) (int, error)
	UpdateBankData(ctx context.Context, bankData storage.BankData) error
	DeleteBankData(ctx context.Context, bankData storage.BankData) error

	AddBinaryData(ctx context.Context, binaryData storage.BinaryData) (int, error)
	UpdateBinaryData(ctx context.Context, binaryData storage.BinaryData) error
	DeleteBinaryData(ctx context.Context, binaryData storage.BinaryData) error

	GetAllRecords(ctx context.Context) (*storage.RecordsList, error)
}
