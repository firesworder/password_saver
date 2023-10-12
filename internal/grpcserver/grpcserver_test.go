package grpcserver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firesworder/password_saver/internal/mocks"
	pb "github.com/firesworder/password_saver/proto"
)

func NewTestGRPCServer(t *testing.T) *GRPCService {
	s := &mocks.Server{}
	grpcs, err := NewGRPCService(s)
	require.NoError(t, err)
	return grpcs
}

func TestNewGRPCServer(t *testing.T) {
	s := &mocks.Server{}
	grpcs, err := NewGRPCService(s)
	require.NoError(t, err)
	require.NotEmpty(t, grpcs)
}

func TestGRPCServer_RegisterUser(t *testing.T) {
	gs := NewTestGRPCServer(t)

	type args struct {
		login, password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			args:    args{login: "log", password: "pass"},
			wantErr: false,
		},
		{
			name:    "Test 2. Error input.",
			args:    args{login: "admin", password: "admin"},
			wantErr: true,
		},
		{
			name:    "Test 3. Server error.",
			args:    args{login: "log"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gs.RegisterUser(context.Background(),
				&pb.RegisterUserRequest{Login: tt.args.login, Password: tt.args.password})
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEqual(t, "", resp.Token)
			}
		})
	}
}

func TestGRPCServer_LoginUser(t *testing.T) {
	gs := NewTestGRPCServer(t)

	type args struct {
		login, password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			args:    args{login: "log", password: "pass"},
			wantErr: false,
		},
		{
			name:    "Test 2. Error input.",
			args:    args{login: "admin", password: "admin"},
			wantErr: true,
		},
		{
			name:    "Test 3. Server error.",
			args:    args{login: "log"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gs.LoginUser(context.Background(),
				&pb.LoginUserRequest{Login: tt.args.login, Password: tt.args.password})
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEqual(t, "", resp.Token)
			}
		})
	}
}

func TestGRPCServer_AddTextDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.TextData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.TextData{TextData: "test data", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error.",
			input:   pb.TextData{TextData: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gs.AddTextDataRecord(context.Background(), &pb.AddTextDataRequest{TextData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEqual(t, 0, resp.Id)
			}
		})
	}
}

func TestGRPCServer_UpdateTextDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.TextData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.TextData{Id: 50, TextData: "test data", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error on id = 100).",
			input:   pb.TextData{Id: 100, TextData: "test data"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.UpdateTextDataRecord(context.Background(), &pb.UpdateTextDataRequest{TextData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_DeleteTextDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			id:      50,
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error on id = 100).",
			id:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.DeleteTextDataRecord(context.Background(), &pb.DeleteTextDataRequest{Id: tt.id})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_AddBankDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.BankData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.BankData{CardNumber: "1122334455667788", CardExpiry: "22/22", Cvv: "334", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error.",
			input:   pb.BankData{CardNumber: "1122334455667788", CardExpiry: "22/22", Cvv: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gs.AddBankDataRecord(context.Background(), &pb.AddBankDataRequest{BankData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEqual(t, 0, resp.Id)
			}
		})
	}
}

func TestGRPCServer_UpdateBankDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.BankData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.BankData{CardNumber: "1122334455667788", CardExpiry: "22/22", Cvv: "334", MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error if ID=100).",
			input:   pb.BankData{Id: 100, CardNumber: "", CardExpiry: "", Cvv: "334", MetaInfo: "mi1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.UpdateBankDataRecord(context.Background(), &pb.UpdateBankDataRequest{BankData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_DeleteBankDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			id:      50,
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error on id = 100).",
			id:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.DeleteBankDataRecord(context.Background(), &pb.DeleteBankDataRequest{Id: tt.id})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_AddBinaryDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.BinaryData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.BinaryData{BinaryData: []byte("test data"), MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error.",
			input:   pb.BinaryData{BinaryData: nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gs.AddBinaryDataRecord(context.Background(), &pb.AddBinaryDataRequest{BinaryData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.NotEqual(t, 0, resp.Id)
			}
		})
	}
}

func TestGRPCServer_UpdateBinaryDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		input   pb.BinaryData
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			input:   pb.BinaryData{BinaryData: []byte("test data"), MetaInfo: "mi1"},
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error if ID=100).",
			input:   pb.BinaryData{Id: 100, BinaryData: nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.UpdateBinaryDataRecord(context.Background(), &pb.UpdateBinaryDataRequest{BinaryData: &tt.input})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_DeleteBinaryDataRecord(t *testing.T) {
	gs := NewTestGRPCServer(t)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Test 1. Correct request.",
			id:      50,
			wantErr: false,
		},
		{
			name:    "Test 2. Server error(special error on id = 100).",
			id:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gs.DeleteBinaryDataRecord(context.Background(), &pb.DeleteBinaryDataRequest{Id: tt.id})
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGRPCServer_GetAllRecords(t *testing.T) {
	gs := NewTestGRPCServer(t)

	resp, err := gs.GetAllRecords(context.Background(), &pb.GetAllRecordsRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, resp.TextDataList)
	require.NotEmpty(t, resp.BankDataList)
	require.NotEmpty(t, resp.BinaryDataList)
}
