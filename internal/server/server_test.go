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

func Test_generateRandom(t *testing.T) {
	bytesLen := 32

	randBytes, err := generateRandom(bytesLen)
	require.NoError(t, err)
	assert.Equal(t, bytesLen, len(randBytes))
}

func Test_generateToken(t *testing.T) {
	token, err := generateToken()
	require.NoError(t, err)
	assert.Equal(t, 32, len(token))
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

func TestServer_AddTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	err = s.AddTextData(ctx, td)
	require.NoError(t, err)
}

func TestServer_UpdateTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	err = s.AddTextData(ctx, td)
	require.NoError(t, err)

	tdUpdate := storage.TextData{ID: 1, TextData: "Definitely not Neeko!"}
	err = s.UpdateTextData(ctx, tdUpdate)
	require.NoError(t, err)
}

func TestServer_DeleteTextData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	td := storage.TextData{TextData: "I will...Neeko!"}
	err = s.AddTextData(ctx, td)
	require.NoError(t, err)

	tdToDelete := storage.TextData{ID: 1}
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
	err = s.AddBankData(ctx, bankD)
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
	err = s.AddBankData(ctx, bankD)
	require.NoError(t, err)

	bankDUpdate := storage.BankData{
		ID:         1,
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
	err = s.AddBankData(ctx, bankD)
	require.NoError(t, err)

	bankDDelete := storage.BankData{ID: 1}
	err = s.DeleteBankData(ctx, bankDDelete)
	require.NoError(t, err)
}

func TestServer_AddBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)
}

func TestServer_UpdateBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)

	binDUpdate := storage.BinaryData{ID: 1, BinaryData: []byte("Cipher")}
	err = s.UpdateBinaryData(ctx, binDUpdate)
	require.NoError(t, err)
}

func TestServer_DeleteBinaryData(t *testing.T) {
	ctx := context.Background()
	s, err := NewServer()
	require.NoError(t, err)

	binD := storage.BinaryData{BinaryData: []byte("Bayonetta")}
	err = s.AddBinaryData(ctx, binD)
	require.NoError(t, err)

	binDDelete := storage.BinaryData{ID: 1}
	err = s.DeleteBinaryData(ctx, binDDelete)
	require.NoError(t, err)
}
