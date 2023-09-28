package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func devStorage(t *testing.T) *Storage {
	s, err := NewStorage(storage.DevDSN)
	if err = s.conn.Ping(); err != nil {
		t.Skip("dev bd is not available, skipping")
	}
	return s
}

func clearTables(t *testing.T, db *sql.DB) {
	var err error

	_, err = db.ExecContext(context.Background(), "DELETE FROM records")
	require.NoError(t, err)

	_, err = db.ExecContext(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
}

func TestNewStorage(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		wantErr bool
	}{
		{
			name:    "Test 1. DSN is correct and available(dev only!)",
			dsn:     storage.DevDSN,
			wantErr: false,
		},
		{
			name:    "Test 2. DSN is correct, but is not available",
			dsn:     "postgresql://postgres:admin@localhost:0000/password_saver",
			wantErr: true,
		},
		{
			name:    "Test 3. Incorrect DSN",
			dsn:     "demoDSN",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStorage(tt.dsn)
			if err != nil && tt.dsn == storage.DevDSN {
				t.Skip("devDSN seems not be available")
			}
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
