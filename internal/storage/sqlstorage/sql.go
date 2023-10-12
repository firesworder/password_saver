package sqlstorage

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/firesworder/password_saver/internal/storage"
)

// Storage основной тип пакета.
// В себе хранит подключение к БД, а также ссылки на каждый из репозиториев данных(размещенных в этой БД).
type Storage struct {
	conn *sql.DB

	UserRep   storage.UserRepository
	RecordRep storage.RecordRepository
}

// NewStorage конструктор sqlstorage, осуществляющий подключение к БД и создание необходимых таблиц.
func NewStorage(DSN string) (*Storage, error) {
	// Этот метод вызывается при инициализации сервера, поэтому использую общий контекст
	ctx := context.Background()
	var err error
	var conn *sql.DB
	if conn, err = sql.Open("pgx", DSN); err != nil {
		return nil, err
	}
	db := Storage{conn: conn}

	if _, err = conn.ExecContext(ctx, createTablesSQL); err != nil {
		return nil, err
	}
	db.UserRep = &User{conn: conn}

	db.RecordRep = &RecordRepository{conn: conn}
	return &db, nil
}
