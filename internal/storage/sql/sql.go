package sql

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
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

	db := Storage{}
	err := db.openDBConnection(DSN)
	if err != nil {
		return nil, err
	}
	err = db.createTablesIfNotExist(ctx)
	if err != nil {
		return nil, err
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
