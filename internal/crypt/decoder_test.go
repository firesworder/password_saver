package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDecoder(t *testing.T) {
	tests := []struct {
		name    string
		fp      string
		wantErr bool
	}{
		{
			name:    "Test 1. Public key correct.",
			fp:      "tests/privKey_test.pem",
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
			enc, err := NewDecoder(tt.fp)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.NotEmpty(t, enc)
			}
		})
	}
}

func TestEncodeDecode(t *testing.T) {
	correctCFP := "tests/cert_test.pem"
	enc, err := NewEncoder(correctCFP)
	require.NoError(t, err)

	correctPKFP := "tests/privKey_test.pem"
	dec, err := NewDecoder(correctPKFP)
	require.NoError(t, err)

	demoMsg := "demo_message"
	// correct call
	encMsg, err := enc.Encode([]byte(demoMsg))
	require.NoError(t, err)
	msg, err := dec.Decode(encMsg)
	require.NoError(t, err)

	assert.Equal(t, demoMsg, string(msg))
}
