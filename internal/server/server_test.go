package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/mocks/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name      string
		wantError bool
	}{
		{
			name:      "Test 1. Basic test",
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewServer()
			assert.Equal(t, tt.wantError, err != nil)
			if err != nil {
				assert.NotEmpty(t, s.env)
			}
		})
	}
}

func TestServer_RegisterUser(t *testing.T) {
	s, err := NewServer()
	require.NoError(t, err)
	ctx := context.Background()

	tests := []struct {
		name          string
		user          storage.User
		uRepState     storage.UserRepository
		wantURepState storage.UserRepository
		wantError     bool
	}{
		{
			name: "Test 1. Successful user reg",
			user: storage.User{Login: "Ayaka", HashedPassword: "hashed_pass1"},
			uRepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
			}},
			wantURepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
				"Ayaka": {Login: "Ayaka", HashedPassword: "hashed_pass1"},
			}},
			wantError: false,
		},
		{
			name: "Test 2. Error user reg, user exists",
			user: storage.User{Login: "Ayato", HashedPassword: "hashed_pass1"},
			uRepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
			}},
			wantURepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
			}},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.uRep = tt.uRepState

			gotErr := s.RegisterUser(ctx, tt.user)
			assert.Equal(t, tt.wantError, gotErr != nil)
			assert.Equal(t, tt.wantURepState, s.uRep)
		})
	}
}

func TestServer_LoginUser(t *testing.T) {
	s, err := NewServer()
	require.NoError(t, err)
	ctx := context.Background()

	tests := []struct {
		name      string
		user      storage.User
		uRepState storage.UserRepository
		wantError bool
	}{
		{
			name: "Test 1. Successful user login",
			user: storage.User{Login: "Ayato", HashedPassword: "hashed_pass2"},
			uRepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
				"Ayaka": {Login: "Ayaka", HashedPassword: "hashed_pass1"},
			}},
			wantError: false,
		},
		{
			name: "Test 2. Error user log, user not exists",
			user: storage.User{Login: "Ayaka", HashedPassword: "hashed_pass1"},
			uRepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
			}},
			wantError: true,
		},
		{
			name: "Test 2. Error user log, user pass not correct",
			user: storage.User{Login: "Ayaka", HashedPassword: "random_pass"},
			uRepState: &users.MockUser{Users: map[string]storage.User{
				"Ayato": {Login: "Ayato", HashedPassword: "hashed_pass2"},
				"Ayaka": {Login: "Ayaka", HashedPassword: "hashed_pass1"},
			}},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.uRep = tt.uRepState

			gotErr := s.LoginUser(ctx, tt.user)
			assert.Equal(t, tt.wantError, gotErr != nil)
		})
	}
}
