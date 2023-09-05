package sqlstorage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const devDSN = "postgresql://postgres:admin@localhost:5432/password_saver"

func TestCommon(t *testing.T) {
	stor, err := NewStorage(devDSN)
	assert.NoError(t, err)
	require.NotEmpty(t, stor)
	assert.NotEmpty(t, stor.Connection)
}
