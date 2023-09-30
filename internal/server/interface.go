package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
)

// IServer интерфейс к Server(основное назначение - подмена моком, при тестировании)
type IServer interface {
	RegisterUser(ctx context.Context, user storage.User) (string, error)
	LoginUser(ctx context.Context, user storage.User) (string, error)

	AddRecord(ctx context.Context, rawRecord interface{}) (int, error)
	UpdateRecord(ctx context.Context, rawRecord interface{}) error
	DeleteRecord(ctx context.Context, rawRecord interface{}) error

	GetAllRecords(ctx context.Context) (*storage.RecordsList, error)
}
