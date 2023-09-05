package repositories

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	conn := getConnection(t)
	defer clearUserTable(t, conn)
	uRep := User{Conn: conn}

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
}

func TestUser_GetUser(t *testing.T) {
	conn := getConnection(t)
	defer clearUserTable(t, conn)
	uRep := User{Conn: conn}

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
}
