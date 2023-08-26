package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name      string
		wantError bool
	}{
		{
			name:      "Test 1. Basic test",
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewServer()
			assert.Equal(t, tt.wantError, err != nil)
			if err != nil {
				assert.NotEmpty(t, s.env)
			}
		})
	}
}
