package grpcserver

import (
	"github.com/firesworder/password_saver/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGRPCServer(t *testing.T) {
	s, err := server.NewServer()
	require.NoError(t, err)

	tests := []struct {
		name      string
		s         *server.Server
		wantError bool
	}{
		{
			name:      "Test 1. Base test",
			s:         s,
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpcS, err := NewGRPCServer(tt.s)
			assert.Equal(t, tt.wantError, err != nil)
			if err != nil {
				assert.NotEmpty(t, grpcS)
				assert.NotEmpty(t, grpcS.s)
			}
		})
	}
}
