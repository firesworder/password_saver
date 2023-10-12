package agentstate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/firesworder/password_saver/internal/storage"
)

func TestNewState(t *testing.T) {
	state := NewState()
	assert.NotEmpty(t, state)
}

func TestState_Get(t *testing.T) {
	state := NewState()
	td := storage.TextData{ID: 1, TextData: "test_1", MetaInfo: "mi_test_1"}
	bankd := storage.BankData{
		ID: 10, CardNumber: "0011223344556677", CardExpire: "09/23", CVV: "234", MetaInfo: "mi_test_10",
	}
	bind := storage.BinaryData{ID: 20, BinaryData: []byte("test_20"), MetaInfo: "mi_test_20"}
	state.textDL[1], state.bankDL[10], state.binaryDL[20] = td, bankd, bind

	type searchedRecord struct {
		id       int
		dataType string
	}

	tests := []struct {
		name           string
		searchedRecord searchedRecord
		wantResult     interface{}
		wantErr        error
	}{
		{
			name:           "Test 1. Text data",
			searchedRecord: searchedRecord{id: 1, dataType: "text"},
			wantResult:     td,
			wantErr:        nil,
		},
		{
			name:           "Test 2. Bank data",
			searchedRecord: searchedRecord{id: 10, dataType: "bank"},
			wantResult:     bankd,
			wantErr:        nil,
		},
		{
			name:           "Test 3. Binary data",
			searchedRecord: searchedRecord{id: 20, dataType: "binary"},
			wantResult:     bind,
			wantErr:        nil,
		},
		{
			name:           "Test 4. Unknown datatype",
			searchedRecord: searchedRecord{id: 10, dataType: "prank"},
			wantResult:     nil,
			wantErr:        fmt.Errorf("record was not found"),
		},
		{
			name:           "Test 5. Unknown ID",
			searchedRecord: searchedRecord{id: 50, dataType: "bank"},
			wantResult:     nil,
			wantErr:        fmt.Errorf("record was not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, gotErr := state.Get(tt.searchedRecord.id, tt.searchedRecord.dataType)
			assert.Equal(t, tt.wantResult, v)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestState_Set(t *testing.T) {
	state := NewState()

	tests := []struct {
		name   string
		record interface{}
	}{
		{
			name:   "Test 1. Text data",
			record: storage.TextData{ID: 1, TextData: "test_1", MetaInfo: "mi_test_1"},
		},
		{
			name: "Test 2. Bank data",
			record: storage.BankData{
				ID: 10, CardNumber: "0011223344556677", CardExpire: "09/23", CVV: "234", MetaInfo: "mi_test_10",
			},
		},
		{
			name:   "Test 3. Binary data",
			record: storage.BinaryData{ID: 20, BinaryData: []byte("test_20"), MetaInfo: "mi_test_20"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state.Set(tt.record)
			switch r := tt.record.(type) {
			case storage.TextData:
				assert.Contains(t, state.textDL, r.ID)
			case storage.BankData:
				assert.Contains(t, state.bankDL, r.ID)
			case storage.BinaryData:
				assert.Contains(t, state.binaryDL, r.ID)
			}
		})
	}
}

func TestState_Delete(t *testing.T) {
	state := NewState()
	td := storage.TextData{ID: 1, TextData: "test_1", MetaInfo: "mi_test_1"}
	bankd := storage.BankData{
		ID: 10, CardNumber: "0011223344556677", CardExpire: "09/23", CVV: "234", MetaInfo: "mi_test_10",
	}
	bind := storage.BinaryData{ID: 20, BinaryData: []byte("test_20"), MetaInfo: "mi_test_20"}
	state.textDL[1], state.bankDL[10], state.binaryDL[20] = td, bankd, bind

	tests := []struct {
		name     string
		recordID int
		wantErr  error
	}{
		{
			name:     "Test 1. Text data",
			recordID: 1,
			wantErr:  nil,
		},
		{
			name:     "Test 2. Bank data",
			recordID: 10,
			wantErr:  nil,
		},
		{
			name:     "Test 3. Binary data",
			recordID: 20,
			wantErr:  nil,
		},
		{
			name:     "Test 5. Unknown ID",
			recordID: 500,
			wantErr:  fmt.Errorf("record was not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := state.Delete(tt.recordID)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
	assert.Empty(t, state.textDL)
	assert.Empty(t, state.bankDL)
	assert.Empty(t, state.binaryDL)
}
