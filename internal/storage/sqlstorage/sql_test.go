package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func devStorage(t *testing.T) *Storage {
	s := Storage{}
	err := s.openDBConnection(storage.DevDSN)
	require.NoError(t, err)
	if err = s.Connection.Ping(); err != nil {
		t.Skip("dev bd is not available, skipping")
	}

	err = s.createTablesIfNotExist(context.Background())
	require.NoError(t, err)
	return &s
}

func clearTables(t *testing.T, db *sql.DB) {
	var err error

	_, err = db.ExecContext(context.Background(), "DELETE FROM records")
	require.NoError(t, err)

	_, err = db.ExecContext(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
}

func TestNewStorageCorrect(t *testing.T) {
	testEnv := env.Environment{
		DSN:            storage.DevDSN,
		CertFile:       "tests/cert_test.pem",
		PrivateKeyFile: "tests/privKey_test.pem",
	}

	s, err := NewStorage(&testEnv)
	if err != nil && testEnv.DSN == storage.DevDSN {
		t.Skip("devDSN is not available")
	}
	require.NoError(t, err)

	require.NotEmpty(t, s)
	assert.NotEmpty(t, s.UserRep)
	assert.NotEmpty(t, s.TextRep)
	assert.NotEmpty(t, s.BankRep)
	assert.NotEmpty(t, s.BinaryRep)
}

func TestNewStorageErrors(t *testing.T) {
	tests := []struct {
		name string
		env  env.Environment
	}{
		{
			name: "Test 1. DevDSN is not available",
			env: env.Environment{
				DSN:            "demoDSN",
				CertFile:       "tests/cert_test.pem",
				PrivateKeyFile: "tests/privKey_test.pem",
			},
		},
		{
			name: "Test 2. Cert file is not set",
			env: env.Environment{
				DSN:            storage.DevDSN,
				CertFile:       "",
				PrivateKeyFile: "tests/privKey_test.pem",
			},
		},
		{
			name: "Test 3. PrivateKey file is not set",
			env: env.Environment{
				DSN:            storage.DevDSN,
				CertFile:       "tests/cert_test.pem",
				PrivateKeyFile: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStorage(&tt.env)
			require.NotEmpty(t, err)
		})
	}
}
