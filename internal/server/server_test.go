package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"testing"
)

var testUser = storage.User{ID: 100, Login: "user", HashedPassword: "password1"}
var testToken = "user_token_1"

func NewTestServer(t *testing.T) *Server {
	genToken, err := generateRandom(32)
	require.NoError(t, err)

	s := &Server{
		authUsers: map[string]storage.User{"user_token_1": {ID: 100, Login: "user", HashedPassword: "password1"}},
		uRep:      &mocks.UserRepository{},
		tRep:      &mocks.TextDataRepository{},
		bankRep:   &mocks.BankDataRepository{},
		binRep:    &mocks.BinaryDataRepository{},
		genToken:  genToken,
	}
	return s
}

func TestServer_getUserFromContext(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name     string
		md       metadata.MD
		wantUser *storage.User
		wantErr  bool
	}{
		{
			name:     "Test 1. Correct MD",
			md:       metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			wantUser: &testUser,
			wantErr:  false,
		},
		{
			name:     "Test 2. Empty MD",
			md:       nil,
			wantUser: nil,
			wantErr:  true,
		},
		{
			name:     "Test 3. MD without token param",
			md:       metadata.New(map[string]string{"some_param": "some_value"}),
			wantUser: nil,
			wantErr:  true,
		},
		{
			name:     "Test 4. MD with unknown token",
			md:       metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			wantUser: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCtx := context.Background()
			if tt.md != nil {
				inputCtx = metadata.NewIncomingContext(inputCtx, tt.md)
			}

			gotUser, err := s.getUserFromContext(inputCtx)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func TestServer_AddTextData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		td      storage.TextData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{TextData: "td", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			td:      storage.TextData{TextData: "td", MetaInfo: "mi1"},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			id, err := s.AddTextData(ctx, tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, err != nil, id == 0)
		})
	}
}

func TestServer_UpdateTextData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		td      storage.TextData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{ID: 200, TextData: "td", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			td:      storage.TextData{ID: 200, TextData: "td", MetaInfo: "mi1"},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{ID: 100, TextData: "td", MetaInfo: "mi1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.UpdateTextData(ctx, tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_DeleteTextData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		td      storage.TextData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{ID: 200},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			td:      storage.TextData{ID: 200},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			td:      storage.TextData{ID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.DeleteTextData(ctx, tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_AddBankData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BankData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{CardNumber: "0011334466779900", CardExpire: "00/23", CVV: "252"},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BankData{CardNumber: "0011334466779900", CardExpire: "00/23", CVV: "252"},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{CVV: "252"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			id, err := s.AddBankData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, err != nil, id == 0)
		})
	}
}

func TestServer_UpdateBankData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BankData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{ID: 200, CardNumber: "0011334466779900", CardExpire: "00/23", CVV: "252"},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BankData{ID: 200, CardNumber: "0011334466779900", CardExpire: "00/23", CVV: "252"},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{ID: 100, CVV: "252"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.UpdateBankData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_DeleteBankData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BankData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{ID: 200},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BankData{ID: 200},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BankData{ID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.DeleteBankData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_AddBinaryData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BinaryData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{BinaryData: []byte("binary data")},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BinaryData{BinaryData: []byte("binary data")},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			id, err := s.AddBinaryData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, err != nil, id == 0)
		})
	}
}

func TestServer_UpdateBinaryData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BinaryData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{ID: 200, BinaryData: []byte("binary data")},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BinaryData{ID: 200, BinaryData: []byte("binary data")},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{ID: 100, BinaryData: []byte("binary data")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.UpdateBinaryData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_DeleteBinaryData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		bd      storage.BinaryData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{ID: 200},
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			bd:      storage.BinaryData{ID: 200},
			wantErr: true,
		},
		{
			name:    "Test 3. Repository error",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			bd:      storage.BinaryData{ID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.DeleteBinaryData(ctx, tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_GetAllRecords(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name    string
		md      metadata.MD
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: "user_token_1"}),
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user token",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			result, err := s.GetAllRecords(ctx)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEmpty(t, result)
				assert.NotEmpty(t, result.TextDataList)
				assert.NotEmpty(t, result.BankDataList)
				assert.NotEmpty(t, result.BinaryDataList)
			}
		})
	}
}
