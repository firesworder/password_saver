package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_generateRandom(t *testing.T) {
	token, err := generateRandom(32)
	assert.NoError(t, err)
	assert.Equal(t, 32, len(token))
}

func TestServer_generateToken(t *testing.T) {
	s := NewTestServer(t)
	token, err := s.generateToken()
	require.NoError(t, err)
	assert.Equal(t, 32, len(token))
}

func TestServer_RegisterUser(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		u       storage.User
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			u:       storage.User{Login: "uslog", HashedPassword: "uspass"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error",
			u:       storage.User{Login: "demo", HashedPassword: "demo"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := s.RegisterUser(context.Background(), tt.u)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NotEqual(t, "", token)
				assert.Contains(t, s.authUsers, token)
			}
		})
	}
}

func TestServer_LoginUser(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		u       storage.User
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			u:       storage.User{Login: "uslog", HashedPassword: "uspass"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error",
			u:       storage.User{Login: "demo", HashedPassword: "demo"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := s.LoginUser(context.Background(), tt.u)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NotEqual(t, "", token)
				assert.Contains(t, s.authUsers, token)
			}
		})
	}
}
