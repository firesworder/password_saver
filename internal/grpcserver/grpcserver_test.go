package grpcserver

import (
	"context"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/storage/mocks/users"
	pb "github.com/firesworder/password_saver/proto"
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
				assert.NotEmpty(t, grpcS.serv)
			}
		})
	}
}

func TestGRPCServer_RegisterUser(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.RegisterUserRequest
		wantResp *pb.RegisterUserResponse
		wantErr  error
	}{
		{
			name:     "Test 1. Basic test",
			req:      &pb.RegisterUserRequest{Login: "Ayaka", Password: "hashed_pass2"},
			wantResp: &pb.RegisterUserResponse{Token: "token_template_reg"},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.RegisterUser(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_LoginUser(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.LoginUserRequest
		wantResp *pb.LoginUserResponse
		wantErr  error
	}{
		{
			name:     "Test 1. Basic test",
			req:      &pb.LoginUserRequest{Login: "Ayaka", Password: "hashed_pass2"},
			wantResp: nil,
			wantErr:  users.ErrUserNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.LoginUser(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}
