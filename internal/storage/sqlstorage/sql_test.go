package sqlstorage

import (
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommon(t *testing.T) {
	stor, err := NewStorage(storage.DevDSN)
	assert.NoError(t, err)
	require.NotEmpty(t, stor)
	assert.NotEmpty(t, stor.Connection)
	// reps
	assert.NotEmpty(t, stor.UserRep)
	assert.NotEmpty(t, stor.TextRep)
	assert.NotEmpty(t, stor.BankRep)
	assert.NotEmpty(t, stor.BinaryRep)
}
