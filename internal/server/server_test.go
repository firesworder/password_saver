package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
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

func Test_generateRandom(t *testing.T) {
	bytesLen := 32

	randBytes, err := generateRandom(bytesLen)
	require.NoError(t, err)
	assert.Equal(t, bytesLen, len(randBytes))
}

func TestServer_RegisterUser(t *testing.T) {
	s, err := NewServer()
	require.NoError(t, err)
	ctx := context.Background()

	tests := []struct {
		name      string
		user      storage.User
		wantError bool
	}{
		{
			name:      "Test 1. Successful user reg",
			user:      storage.User{Login: "Ayaka", HashedPassword: "hashed_pass1"},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := s.RegisterUser(ctx, tt.user)
			assert.Equal(t, tt.wantError, gotErr != nil)
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
		wantError bool
	}{
		{
			name:      "Test 1. Error user log, user not exists",
			user:      storage.User{Login: "Ayaka", HashedPassword: "hashed_pass1"},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := s.LoginUser(ctx, tt.user)
			require.Equal(t, tt.wantError, gotErr != nil)
		})
	}
}

func TestServer_AddTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	td.ID, err = s.AddTextData(ctx, td)
	require.NoError(t, err)
}

func TestServer_UpdateTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	td.ID, err = s.AddTextData(ctx, td)
	require.NoError(t, err)

	tdUpdate := storage.TextData{ID: td.ID, TextData: "Definitely not Neeko!"}
	err = s.UpdateTextData(ctx, tdUpdate)
	require.NoError(t, err)
}

func TestServer_DeleteTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	td.ID, err = s.AddTextData(ctx, td)
	require.NoError(t, err)

	tdToDelete := storage.TextData{ID: td.ID}
	err = s.DeleteTextData(ctx, tdToDelete)
	require.NoError(t, err)
}

func TestServer_AddBankData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	bankD := storage.BankData{
		CardNumber: "0011 2233 4455 6677",
		CardExpire: "88/00",
		CVV:        "456",
	}
	bankD.ID, err = s.AddBankData(ctx, bankD)
	require.NoError(t, err)
}

func TestServer_UpdateBankData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	bankD := storage.BankData{
		CardNumber: "0011 2233 4455 6677",
		CardExpire: "88/00",
		CVV:        "456",
	}
	bankD.ID, err = s.AddBankData(ctx, bankD)
	require.NoError(t, err)

	bankDUpdate := storage.BankData{
		ID:         bankD.ID,
		CardNumber: "8800 9900 7700 6666",
		CardExpire: "88/00",
		CVV:        "456",
	}
	err = s.UpdateBankData(ctx, bankDUpdate)
	require.NoError(t, err)
}

func TestServer_DeleteBankData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	bankD := storage.BankData{
		CardNumber: "0011 2233 4455 6677",
		CardExpire: "88/00",
		CVV:        "456",
	}
	bankD.ID, err = s.AddBankData(ctx, bankD)
	require.NoError(t, err)

	bankDDelete := storage.BankData{ID: bankD.ID}
	err = s.DeleteBankData(ctx, bankDDelete)
	require.NoError(t, err)
}

func TestServer_AddBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	binD.ID, err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)
}

func TestServer_UpdateBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	binD.ID, err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)

	binDUpdate := storage.BinaryData{ID: binD.ID, BinaryData: []byte("Cipher")}
	err = s.UpdateBinaryData(ctx, binDUpdate)
	require.NoError(t, err)
}

func TestServer_DeleteBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	binD.ID, err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)

	binDDelete := storage.BinaryData{ID: binD.ID}
	err = s.DeleteBinaryData(ctx, binDDelete)
	require.NoError(t, err)
}
