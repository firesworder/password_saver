package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"
)

type User struct {
	Conn *sql.DB
}

func (ur *User) CreateUser(ctx context.Context, u storage.User) (*storage.User, error) {
	var id int

	err := ur.Conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) RETURNING id", u.Login, u.HashedPassword,
	).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, storage.ErrLoginExist
		}
		return nil, err
	}

	return &storage.User{ID: id, Login: u.Login, HashedPassword: u.HashedPassword}, nil
}

func (ur *User) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	gotUser := storage.User{}
	err := ur.Conn.QueryRowContext(ctx,
		"SELECT id, login, password FROM users WHERE login = $1 LIMIT 1", u.Login,
	).Scan(&gotUser.ID, &gotUser.Login, &gotUser.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrLoginNotExist
		}
		return nil, err
	}
	return &gotUser, nil
}
