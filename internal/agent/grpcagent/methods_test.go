package grpcagent

import (
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func getGRPCConn(t *testing.T) (*server.Server, *GRPCAgent, func()) {
	s, err := server.NewServer()
	require.NoError(t, err)
	testGS := startTestGRPCServer(t, s)
	gotGAgent, err := NewGRPCAgent(devTestAddr)
	require.NoError(t, err)

	closer := func() {
		err = gotGAgent.Close()
		require.NoError(t, err)
		testGS.GracefulStop()
	}

	return s, gotGAgent, closer
}

func TestGRPCAgent_RegisterUser(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	type args struct {
		login, password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1. Correct use",
			args:    args{login: "user1", password: "password1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := a.RegisterUser(tt.args.login, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			if err != nil {
				assert.NotEqual(t, token, "")
			}
		})
	}
	closer()
}

func TestGRPCAgent_LoginUser(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	type args struct {
		login, password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1. Correct use",
			args:    args{login: "user1", password: "password1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := a.LoginUser(tt.args.login, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.NotEqual(t, token, "")
			}
		})
	}
	closer()
}

func TestGRPCAgent_CreateBankDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		bd      storage.BankData
		wantErr bool
	}{
		{
			name: "Test 1. Correct use",
			bd: storage.BankData{
				CardNumber: "0011 2233 4455 6677",
				CardExpire: "09/23",
				CVV:        "123",
				MetaInfo:   "meta 1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := a.CreateBankDataRecord(tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_CreateBinaryDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		binD    storage.BinaryData
		wantErr bool
	}{
		{
			name: "Test 1. Correct use",
			binD: storage.BinaryData{
				BinaryData: []byte("Binary info 1"),
				MetaInfo:   "meta 1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := a.CreateBinaryDataRecord(tt.binD)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_CreateTextDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		td      storage.TextData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			td: storage.TextData{
				TextData: "Text data 1",
				MetaInfo: "meta 1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := a.CreateTextDataRecord(tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_UpdateBankDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		bd      storage.BankData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			bd: storage.BankData{
				CardNumber: "1122 3344 5566 7788",
				CardExpire: "12/23",
				CVV:        "123",
				MetaInfo:   "meta 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.UpdateBankDataRecord(tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_UpdateBinaryDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		binD    storage.BinaryData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			binD: storage.BinaryData{
				BinaryData: []byte("Binary info 1"),
				MetaInfo:   "meta 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.UpdateBinaryDataRecord(tt.binD)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_UpdateTextDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		td      storage.TextData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			td: storage.TextData{
				TextData: "Text data 1",
				MetaInfo: "meta 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.UpdateTextDataRecord(tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_DeleteBankDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		bd      storage.BankData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			bd: storage.BankData{
				ID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.DeleteBankDataRecord(tt.bd)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_DeleteBinaryDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		binD    storage.BinaryData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			binD: storage.BinaryData{
				BinaryData: []byte("Binary info 1"),
				MetaInfo:   "meta 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.UpdateBinaryDataRecord(tt.binD)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_DeleteTextDataRecord(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		td      storage.TextData
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			td: storage.TextData{
				TextData: "Text data 1",
				MetaInfo: "meta 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.UpdateTextDataRecord(tt.td)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}

func TestGRPCAgent_ShowAllRecords(t *testing.T) {
	_, a, closer := getGRPCConn(t)

	tests := []struct {
		name    string
		wantRec *storage.RecordsList
		wantErr bool
	}{
		{
			name: "Test 1. Correct user",
			wantRec: &storage.RecordsList{
				TextDataList:   []storage.TextData{},
				BankDataList:   []storage.BankData{},
				BinaryDataList: []storage.BinaryData{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRec, err := a.ShowAllRecords()
			assert.Equal(t, tt.wantRec, gotRec)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
	closer()
}
