package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/crypt"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	Connection *sql.DB

	UserRep   storage.UserRepository
	TextRep   storage.TextDataRepository
	BankRep   storage.BankDataRepository
	BinaryRep storage.BinaryDataRepository
}

// NewStorage конструктор sqlstorage, также инициал. шифрование.
func NewStorage(env *env.Environment) (*Storage, error) {
	// Этот метод вызывается при инициализации сервера, поэтому использую общий контекст
	ctx := context.Background()

	var err error
	db := Storage{}
	if err = db.openDBConnection(env.DSN); err != nil {
		return nil, err
	}
	if err = db.createTablesIfNotExist(ctx); err != nil {
		return nil, err
	}

	encoder, err := crypt.NewEncoder(env.CertFile)
	if err != nil {
		return nil, err
	}
	decoder, err := crypt.NewDecoder(env.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	db = Storage{
		Connection: db.Connection,
		UserRep:    &User{Conn: db.Connection},
		TextRep: &TextData{
			Conn: db.Connection, Encoder: encoder, Decoder: decoder,
		},
		BankRep: &BankData{
			Conn: db.Connection, Encoder: encoder, Decoder: decoder,
		},
		BinaryRep: &BinaryData{
			Conn: db.Connection, Encoder: encoder, Decoder: decoder,
		},
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
