package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var Storage, storageErr = sqlstorage.NewStorage(storage.DevDSN)

func clearUserTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
}

func TestUser_CreateUser(t *testing.T) {
	clearUserTable(t, Storage.Connection)
	if storageErr != nil {
		t.Skipf("test skipped, db connection is not avail: %s", storageErr)
	}
	uRep := User{conn: Storage.Connection}

	tests := []struct {
		name     string
		user     storage.User
		wantUser *storage.User
		wantErr  error
	}{
		{
			name:     "Test 1. Common use - user not exist",
			user:     storage.User{Login: "Ayaka", HashedPassword: "Kamisato"},
			wantUser: &storage.User{ID: 1, Login: "Ayaka", HashedPassword: "Kamisato"},
			wantErr:  nil,
		},
		{
			name:     "Test 2. Common use - user exist",
			user:     storage.User{Login: "Ayaka", HashedPassword: "Kamisato"},
			wantUser: nil,
			wantErr:  storage.ErrLoginExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := uRep.CreateUser(context.Background(), tt.user)
			if tt.wantUser != nil {
				require.NotEmpty(t, gotUser)
				assert.Greater(t, gotUser.ID, 0)
				assert.Equal(t, tt.wantUser.Login, gotUser.Login)
				assert.Equal(t, tt.wantUser.HashedPassword, gotUser.HashedPassword)
			} else {
				assert.Empty(t, gotUser)
			}
			assert.ErrorIs(t, gotErr, tt.wantErr)
		})
	}
	clearUserTable(t, Storage.Connection)
}

func TestUser_GetUser(t *testing.T) {
	if storageErr != nil {
		t.Skipf("test skipped, db connection is not avail: %s", storageErr)
	}

	uRep := User{conn: Storage.Connection}
	newUser, err := uRep.CreateUser(context.Background(), storage.User{Login: "Ayaka", HashedPassword: "Kamisato"})
	require.NoError(t, err)

	tests := []struct {
		name     string
		user     storage.User
		wantUser *storage.User
		wantErr  error
	}{
		{
			name:     "Test 1. Common use - user not exist",
			user:     storage.User{Login: "Ayato"},
			wantUser: nil,
			wantErr:  storage.ErrLoginNotExist,
		},
		{
			name:     "Test 2. Common use - user exist",
			user:     storage.User{Login: "Ayaka", HashedPassword: "Kamisato"},
			wantUser: &storage.User{ID: newUser.ID, Login: "Ayaka", HashedPassword: "Kamisato"},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := uRep.GetUser(context.Background(), tt.user)
			assert.Equal(t, tt.wantUser, gotUser)
			assert.ErrorIs(t, gotErr, tt.wantErr)
		})
	}
	//clearUserTable(t, Storage.Connection)
}
