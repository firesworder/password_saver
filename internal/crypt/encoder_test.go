package crypt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	tests := []struct {
		name    string
		fp      string
		wantErr bool
	}{
		{
			name:    "Test 1. Public key correct.",
			fp:      "tests/cert_test.pem",
			wantErr: false,
		},
		{
			name:    "Test 2. Empty file.",
			fp:      "tests/emptyfile_test.pem",
			wantErr: true,
		},
		{
			name:    "Test 3. File not exist.",
			fp:      "tests/not_exist_file.pem",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc, err := NewEncoder(tt.fp)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.NotEmpty(t, enc)
			}
		})
	}
}

func TestEncoder_Encode(t *testing.T) {
	correctFP := "tests/cert_test.pem"
	enc, err := NewEncoder(correctFP)
	require.NoError(t, err)

	// correct call
	msg, err := enc.Encode([]byte("demo_message"))
	require.NoError(t, err)
	assert.NotEmpty(t, msg)
}
