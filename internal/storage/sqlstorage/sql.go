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

	uRep    storage.UserRepository
	tRep    storage.TextDataRepository
	bankRep storage.BankDataRepository
	binRep  storage.BinaryDataRepository
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
		uRep:       &repositories.User{Conn: db.Connection},
		tRep:       &repositories.TextData{Conn: db.Connection},
		bankRep:    &repositories.BankData{Conn: db.Connection},
		binRep:     &repositories.BinaryData{Conn: db.Connection},
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
