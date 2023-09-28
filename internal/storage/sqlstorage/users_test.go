package sqlstorage

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	ctx := context.Background()

	s := devStorage(t)
	// закрываю соединение с дб
	defer s.conn.Close()
	uRep := User{conn: s.conn}

	// очищаю таблицы перед добавлением новых тестовых данных и по итогам прогона тестов
	clearTables(t, s.conn)
	defer clearTables(t, s.conn)

	// подготовка тестовых данных
	var uID int64
	var err error
	err = s.conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo", "demo").Scan(&uID)
	require.NoError(t, err)

	type args struct {
		login, password string
	}

	tests := []struct {
		name string
		args
		wantErr error
	}{
		{
			name:    "Test 1. Correct new user",
			args:    args{login: "demo1", password: "demo1"},
			wantErr: nil,
		},
		{
			name:    "Test 2. User with that login already exist",
			args:    args{login: "demo", password: "demo"},
			wantErr: storage.ErrLoginExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cUser, err := uRep.CreateUser(ctx, storage.User{Login: tt.login, HashedPassword: tt.password})
			assert.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				assert.NotEqual(t, 0, cUser.ID)
			}
		})
	}
}

func TestUser_GetUser(t *testing.T) {
	ctx := context.Background()

	s := devStorage(t)
	// закрываю соединение с дб
	defer s.conn.Close()
	uRep := User{conn: s.conn}

	// очищаю таблицы перед добавлением новых тестовых данных и по итогам прогона тестов
	clearTables(t, s.conn)
	defer clearTables(t, s.conn)

	// подготовка тестовых данных
	var uID int64
	var err error
	err = s.conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo", "demo").Scan(&uID)
	require.NoError(t, err)

	type args struct {
		login, password string
	}

	tests := []struct {
		name string
		args
		wantErr error
	}{
		{
			name:    "Test 1. Correct login user",
			args:    args{login: "demo", password: "demo"},
			wantErr: nil,
		},
		{
			name:    "Test 2. User not exist",
			args:    args{login: "demo2", password: "demo2"},
			wantErr: storage.ErrLoginNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cUser, err := uRep.GetUser(ctx, storage.User{Login: tt.login, HashedPassword: tt.password})
			assert.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				assert.Equal(t, tt.login, cUser.Login)
				assert.Equal(t, tt.password, cUser.HashedPassword)
			}
		})
	}
}
