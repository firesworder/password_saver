package grpcagent

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGRPCAgent(t *testing.T) {
	// correct grpc agent creation
	ga, err := NewGRPCAgent("localhost:8080", "ca_cert_test.pem")
	require.NoError(t, err)
	assert.NotEmpty(t, ga.grpcClient)
	assert.NotEmpty(t, ga.conn)
	// ga client closing
	err = ga.Close()
	require.NoError(t, err)

	_, err = NewGRPCAgent("localhost:8080", "")
	require.NotEqual(t, nil, err)
}
