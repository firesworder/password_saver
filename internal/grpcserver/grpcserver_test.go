package grpcserver

import (
	"context"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/storage/mocks/bankdata"
	"github.com/firesworder/password_saver/internal/storage/mocks/binarydata"
	"github.com/firesworder/password_saver/internal/storage/mocks/textdata"
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
		name    string
		req     *pb.RegisterUserRequest
		wantErr error
	}{
		{
			name:    "Test 1. Basic test",
			req:     &pb.RegisterUserRequest{Login: "Ayaka", Password: "hashed_pass2"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.RegisterUser(ctx, tt.req)
			assert.ErrorIs(t, tt.wantErr, gotErr)
			if gotErr == nil {
				assert.NotEqual(t, gotResp.Token, "")
			}
		})
	}
}

func TestGRPCServer_LoginUser(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name    string
		req     *pb.LoginUserRequest
		wantErr error
	}{
		{
			name:    "Test 1. Basic test",
			req:     &pb.LoginUserRequest{Login: "Ayaka", Password: "hashed_pass2"},
			wantErr: users.ErrUserNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.LoginUser(ctx, tt.req)
			assert.ErrorIs(t, tt.wantErr, gotErr)
			if gotErr == nil {
				assert.Equal(t, gotResp.Token, "")
			}
		})
	}
}

func TestGRPCServer_AddBankDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.AddBankDataRequest
		wantResp *pb.AddBankDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.AddBankDataRequest{BankData: &pb.BankData{
				CardNumber: "0011 2233 4455 6677",
				CardExpiry: "05/24",
				Cvv:        "987",
				MetaInfo:   "",
			}},
			wantResp: &pb.AddBankDataResponse{Id: 1},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.AddBankDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_AddBinaryDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.AddBinaryDataRequest
		wantResp *pb.AddBinaryDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.AddBinaryDataRequest{BinaryData: &pb.BinaryData{
				BinaryData: []byte("Ayaka"),
				MetaInfo:   "Ayayaka",
			}},
			wantResp: &pb.AddBinaryDataResponse{Id: 1},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.AddBinaryDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_AddTextDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.AddTextDataRequest
		wantResp *pb.AddTextDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.AddTextDataRequest{TextData: &pb.TextData{
				TextData: "Ayato",
				MetaInfo: "Ayayato",
			}},
			wantResp: &pb.AddTextDataResponse{Id: 1},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.AddTextDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_UpdateBankDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.UpdateBankDataRequest
		wantResp *pb.UpdateBankDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.UpdateBankDataRequest{BankData: &pb.BankData{
				Id:         int64(1),
				CardNumber: "0011 2233 4455 6677",
				CardExpiry: "05/24",
				Cvv:        "987",
				MetaInfo:   "",
			}},
			wantResp: nil,
			wantErr:  bankdata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.UpdateBankDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_UpdateBinaryDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.UpdateBinaryDataRequest
		wantResp *pb.UpdateBinaryDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.UpdateBinaryDataRequest{BinaryData: &pb.BinaryData{
				Id:         int64(1),
				BinaryData: []byte("Ayato"),
				MetaInfo:   "",
			}},
			wantResp: nil,
			wantErr:  binarydata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.UpdateBinaryDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_UpdateTextDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.UpdateTextDataRequest
		wantResp *pb.UpdateTextDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.UpdateTextDataRequest{TextData: &pb.TextData{
				Id:       int64(1),
				TextData: "Bayonetta",
				MetaInfo: "a computer game",
			}},
			wantResp: nil,
			wantErr:  textdata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.UpdateTextDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_DeleteBankDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.DeleteBankDataRequest
		wantResp *pb.DeleteBankDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.DeleteBankDataRequest{
				Id: int64(1),
			},
			wantResp: nil,
			wantErr:  bankdata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.DeleteBankDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_DeleteBinaryDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.DeleteBinaryDataRequest
		wantResp *pb.DeleteBinaryDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.DeleteBinaryDataRequest{
				Id: int64(1),
			},
			wantResp: nil,
			wantErr:  binarydata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.DeleteBinaryDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_DeleteTextDataRecord(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.DeleteTextDataRequest
		wantResp *pb.DeleteTextDataResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req: &pb.DeleteTextDataRequest{
				Id: int64(1),
			},
			wantResp: nil,
			wantErr:  textdata.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.DeleteTextDataRecord(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}

func TestGRPCServer_GetAllRecords(t *testing.T) {
	grpcS := GRPCServer{serv: nil}
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *pb.GetAllRecordsRequest
		wantResp *pb.GetAllRecordsResponse
		wantErr  error
	}{
		{
			name: "Test 1. Basic test",
			req:  &pb.GetAllRecordsRequest{},
			wantResp: &pb.GetAllRecordsResponse{
				TextDataList:   []*pb.TextData{},
				BankDataList:   []*pb.BankData{},
				BinaryDataList: []*pb.BinaryData{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.NewServer()
			require.NoError(t, err)
			grpcS.serv = s

			gotResp, gotErr := grpcS.GetAllRecords(ctx, tt.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.ErrorIs(t, tt.wantErr, gotErr)
		})
	}
}
