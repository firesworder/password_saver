package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"testing"
)

var testUser = storage.User{ID: 100, Login: "user", HashedPassword: "password1"}
var testToken = "user_token_1"

// на данном пользователе в моке sqlstorage триггерится тестовая ошибка
var testErrToken = "user_error"
var errUser = storage.User{ID: -1, Login: "err", HashedPassword: "err"}

func NewTestServer(t *testing.T) *Server {
	testEnv := &env.Environment{
		DSN:            storage.DevDSN,
		CertFile:       "test_files/cert.pem",
		PrivateKeyFile: "test_files/privKey.pem",
	}
	s, _ := NewServer(testEnv)
	s.ssql = &sqlstorage.Storage{
		UserRep:   mocks.NewUR(),
		RecordRep: mocks.NewRR(),
	}
	s.authUsers.Store(testToken, testUser)
	s.authUsers.Store(testErrToken, errUser)
	return s
}

func TestNewServer(t *testing.T) {
	tests := []struct {
		name    string
		env     *env.Environment
		wantErr bool
	}{
		{
			name: "Test 1. Correct creation with devDNS",
			env: &env.Environment{
				DSN:            storage.DevDSN,
				CertFile:       "test_files/cert.pem",
				PrivateKeyFile: "test_files/privKey.pem",
			},
			wantErr: false,
		},
		{
			name:    "Test 2. Error, empty env",
			env:     nil,
			wantErr: true,
		},
		{
			name: "Test 3. Error, incorrect devDSN",
			env: &env.Environment{
				DSN:            "demoEnv",
				CertFile:       "test_files/cert.pem",
				PrivateKeyFile: "test_files/privKey.pem",
			},
			wantErr: true,
		},
		{
			name: "Test 4. Error, incorrect cert filepath",
			env: &env.Environment{
				DSN:            "demoEnv",
				CertFile:       "test_files/not_exist.pem",
				PrivateKeyFile: "test_files/privKey.pem",
			},
			wantErr: true,
		},
		{
			name: "Test 5. Error, incorrect priv key filepath",
			env: &env.Environment{
				DSN:            "demoEnv",
				CertFile:       "test_files/cert.pem",
				PrivateKeyFile: "test_files/not_exist.pem",
			},
			wantErr: true,
		},
		{
			name: "Test 6. Error, ",
			env: &env.Environment{
				DSN:            "demoEnv",
				CertFile:       "test_files/cert.pem",
				PrivateKeyFile: "test_files/privKey.pem",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewServer(tt.env)
			if err != nil && tt.env != nil && tt.env.DSN == storage.DevDSN {
				t.Skip("devDSN is not available, skipping")
			}
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
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
			md:       metadata.New(map[string]string{ctxTokenParam: testToken}),
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

func TestServer_getRecordFromData(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name       string
		rawRecord  interface{}
		wantRecord *storage.Record
		wantErr    bool
	}{
		{
			name:       "Test 1. Text data",
			rawRecord:  storage.TextData{ID: 50, TextData: "text data", MetaInfo: "MI", UserID: 100},
			wantRecord: &storage.Record{ID: 50, RecordType: "text", Content: []byte("text data"), MetaInfo: "MI"},
			wantErr:    false,
		},
		{
			name: "Test 2. Bank data",
			rawRecord: storage.BankData{ID: 50, CardNumber: "0011 2233 4455 6677", CardExpire: "09/23", CVV: "333",
				MetaInfo: "MI", UserID: 100},
			wantRecord: &storage.Record{ID: 50, RecordType: "bank", Content: []byte("0011 2233 4455 6677,09/23,333"),
				MetaInfo: "MI"},
			wantErr: false,
		},
		{
			name:       "Test 3. Binary data",
			rawRecord:  storage.BinaryData{ID: 50, BinaryData: []byte("binary data"), MetaInfo: "MI", UserID: 100},
			wantRecord: &storage.Record{ID: 50, RecordType: "binary", Content: []byte("binary data"), MetaInfo: "MI"},
			wantErr:    false,
		},
		{
			name:       "Test 4. Unknown data type",
			rawRecord:  100,
			wantRecord: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := s.getRecordFromData(tt.rawRecord)
			if gotRecord != nil {
				gotRecord.Content, err = s.decoder.Decode(gotRecord.Content)
				require.NoError(t, err)
			}
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantRecord, gotRecord)
		})
	}
}

func TestServer_AddRecord(t *testing.T) {
	s := NewTestServer(t)

	tests := []struct {
		name       string
		md         metadata.MD
		dataRecord interface{}
		wantErr    bool
	}{
		{
			name:       "Test 1. Correct request",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: storage.TextData{TextData: "textData", MetaInfo: "MI1"},
			wantErr:    false,
		},
		{
			name:       "Test 2. Unknown user",
			md:         metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			dataRecord: storage.TextData{TextData: "textData", MetaInfo: "MI1"},
			wantErr:    true,
		},
		{
			name:       "Test 3. Unknown(nil) data type",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: nil,
			wantErr:    true,
		},
		{
			name:       "Test 4. Trigger storage error",
			md:         metadata.New(map[string]string{ctxTokenParam: testErrToken}),
			dataRecord: storage.TextData{TextData: "textData", MetaInfo: "MI1"},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			_, err := s.AddRecord(ctx, tt.dataRecord)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_UpdateRecord(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{ctxTokenParam: testToken}))
	s := NewTestServer(t)

	rid, err := s.AddRecord(ctx, storage.TextData{TextData: "old", MetaInfo: "MI"})
	require.NoError(t, err)
	_, err = s.AddRecord(ctx, storage.TextData{TextData: "old td 2", MetaInfo: "MI2"})
	require.NoError(t, err)

	tests := []struct {
		name       string
		md         metadata.MD
		dataRecord interface{}
		wantErr    bool
	}{
		{
			name:       "Test 1. Correct request",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: storage.TextData{ID: rid, TextData: "textData", MetaInfo: "MI1"},
			wantErr:    false,
		},
		{
			name:       "Test 2. Unknown user",
			md:         metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			dataRecord: storage.TextData{ID: 150, TextData: "textData", MetaInfo: "MI1"},
			wantErr:    true,
		},
		{
			name:       "Test 3. Unknown(nil) data type",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: nil,
			wantErr:    true,
		},
		{
			name:       "Test 4. Trigger storage error",
			md:         metadata.New(map[string]string{ctxTokenParam: testErrToken}),
			dataRecord: storage.TextData{ID: 150, TextData: "textData", MetaInfo: "MI1"},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.UpdateRecord(ctx, tt.dataRecord)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_DeleteRecord(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{ctxTokenParam: testToken}))
	s := NewTestServer(t)

	rid, err := s.AddRecord(ctx, storage.TextData{TextData: "old", MetaInfo: "MI"})
	require.NoError(t, err)
	_, err = s.AddRecord(ctx, storage.TextData{TextData: "old td 2", MetaInfo: "MI2"})
	require.NoError(t, err)

	tests := []struct {
		name       string
		md         metadata.MD
		dataRecord interface{}
		wantErr    bool
	}{
		{
			name:       "Test 1. Correct request",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: storage.TextData{ID: rid},
			wantErr:    false,
		},
		{
			name:       "Test 2. Unknown user",
			md:         metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			dataRecord: storage.TextData{ID: 150},
			wantErr:    true,
		},
		{
			name:       "Test 3. Unknown(nil) data type",
			md:         metadata.New(map[string]string{ctxTokenParam: testToken}),
			dataRecord: nil,
			wantErr:    true,
		},
		{
			name:       "Test 4. Trigger storage error",
			md:         metadata.New(map[string]string{ctxTokenParam: testErrToken}),
			dataRecord: storage.TextData{ID: 150},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			err := s.DeleteRecord(ctx, tt.dataRecord)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestServer_GetAllRecords(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{ctxTokenParam: testToken}))
	s := NewTestServer(t)

	var err error
	_, err = s.AddRecord(ctx, storage.TextData{TextData: "text data 1", MetaInfo: "MI1"})
	require.NoError(t, err)
	_, err = s.AddRecord(ctx, storage.TextData{TextData: "text data 2", MetaInfo: "MI2"})
	require.NoError(t, err)
	_, err = s.AddRecord(ctx, storage.BankData{CardNumber: "0011", CardExpire: "09/23", CVV: "123", MetaInfo: "MI1"})
	require.NoError(t, err)
	_, err = s.AddRecord(ctx, storage.BinaryData{BinaryData: []byte("binary data"), MetaInfo: "MI1"})
	require.NoError(t, err)

	tests := []struct {
		name    string
		md      metadata.MD
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request",
			md:      metadata.New(map[string]string{ctxTokenParam: testToken}),
			wantErr: false,
		},
		{
			name:    "Test 2. Unknown user",
			md:      metadata.New(map[string]string{ctxTokenParam: "unknown_token"}),
			wantErr: true,
		},
		{
			name:    "Test 3. Trigger storage error",
			md:      metadata.New(map[string]string{ctxTokenParam: testErrToken}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = metadata.NewIncomingContext(context.Background(), tt.md)
			records, err := s.GetAllRecords(ctx)
			assert.Equal(t, tt.wantErr, err != nil)

			if records != nil {
				assert.Len(t, records.TextDataList, 2)
				assert.Len(t, records.BankDataList, 1)
				assert.Len(t, records.BinaryDataList, 1)
			}
		})
	}
}
