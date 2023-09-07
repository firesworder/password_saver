package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage/repositories"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	Connection *sql.DB

	UserRep   storage.UserRepository
	TextRep   storage.TextDataRepository
	BankRep   storage.BankDataRepository
	BinaryRep storage.BinaryDataRepository
}

func NewStorage(DSN string) (*Storage, error) {
	// Этот метод вызывается при инициализации сервера, поэтому использую общий контекст
	ctx := context.Background()

	var err error
	db := Storage{}
	if err = db.openDBConnection(DSN); err != nil {
		return nil, err
	}
	if err = db.createTablesIfNotExist(ctx); err != nil {
		return nil, err
	}

	db = Storage{
		Connection: db.Connection,
		UserRep:    &repositories.User{Conn: db.Connection},
		TextRep:    &repositories.TextData{Conn: db.Connection},
		BankRep:    &repositories.BankData{Conn: db.Connection},
		BinaryRep:  &repositories.BinaryData{Conn: db.Connection},
	}
	return &db, nil
}

func (db *Storage) openDBConnection(DSN string) error {
	var err error
	db.Connection, err = sql.Open("pgx", DSN)
	if err != nil {
		return err
	}
	return nil
}

func (db *Storage) createTablesIfNotExist(ctx context.Context) error {
	_, err := db.Connection.ExecContext(ctx, createTablesSQL)
	if err != nil {
		return err
	}
	return nil
}
