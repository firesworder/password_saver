package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"testing"
)

func getConnection(t *testing.T) *sql.DB {
	conn, err := sql.Open("pgx", storage.DevDSN)
	if err != nil {
		t.Skipf("DevDSN is not available")
	}
	return conn
}

func clearUserTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
}

func clearTextDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM textdata")
	require.NoError(t, err)
}

func clearBankDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM bankdata")
	require.NoError(t, err)
}

func clearBinaryDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM binarydata")
	require.NoError(t, err)
}
