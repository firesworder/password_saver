package bankdata

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

var binaryDataState = map[int]storage.BinaryData{
	1: {ID: 1, BinaryData: []byte("Ayaka"), MetaInfo: "1 record"},
	3: {ID: 3, BinaryData: []byte("Ayato")},
}
var lastUsedID = 3

func getStateMap(src map[int]storage.BinaryData) map[int]storage.BinaryData {
	r := map[int]storage.BinaryData{}
	for key, value := range src {
		r[key] = value
	}
	return r
}

func TestMockBinaryData_AddBinaryData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		bd             storage.BinaryData
		wantState      map[int]storage.BinaryData
		wantLastUsedID int
		wantError      error
	}{
		{
			name: "Test 1. Correct binary data",
			bd:   storage.BinaryData{BinaryData: []byte("Bayonetta")},
			wantState: map[int]storage.BinaryData{
				1: {ID: 1, BinaryData: []byte("Ayaka"), MetaInfo: "1 record"},
				3: {ID: 3, BinaryData: []byte("Ayato")},
				4: {ID: 4, BinaryData: []byte("Bayonetta")},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name: "Test 2. Correct binary data, with metaInfo",
			bd:   storage.BinaryData{BinaryData: []byte("Bayonetta"), MetaInfo: "another record"},
			wantState: map[int]storage.BinaryData{
				1: {ID: 1, BinaryData: []byte("Ayaka"), MetaInfo: "1 record"},
				3: {ID: 3, BinaryData: []byte("Ayato")},
				4: {ID: 4, BinaryData: []byte("Bayonetta"), MetaInfo: "another record"},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name: "Test 3. Empty binary data",
			bd:   storage.BinaryData{},
			wantState: map[int]storage.BinaryData{
				1: {ID: 1, BinaryData: []byte("Ayaka"), MetaInfo: "1 record"},
				3: {ID: 3, BinaryData: []byte("Ayato")},
			},
			wantLastUsedID: 3,
			wantError:      ErrEmptyData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBinaryData{BinaryData: getStateMap(binaryDataState), LastUsedID: lastUsedID}

			gotError := rep.AddBinaryData(ctx, tt.bd)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BinaryData)
			assert.Equal(t, tt.wantLastUsedID, rep.LastUsedID)
		})
	}
}

func TestMockBinaryData_UpdateBinaryData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		td        storage.BinaryData
		wantState map[int]storage.BinaryData
		wantError error
	}{
		{
			name: "Test 1. Update text data",
			td:   storage.BinaryData{ID: 1, BinaryData: []byte("Raiden Shogun")},
			wantState: map[int]storage.BinaryData{
				1: {ID: 1, BinaryData: []byte("Raiden Shogun")},
				3: {ID: 3, BinaryData: []byte("Ayato")},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Update text data, ID not set(id=0)",
			td:        storage.BinaryData{BinaryData: []byte("Raiden Shogun")},
			wantState: getStateMap(binaryDataState),
			wantError: ErrNotFound,
		},
		{
			name:      "Test 3. Update text data, empty data",
			td:        storage.BinaryData{ID: 3, BinaryData: []byte{}},
			wantState: getStateMap(binaryDataState),
			wantError: ErrEmptyData,
		},
		{
			name:      "Test 4. Update text data, only metaInfo",
			td:        storage.BinaryData{ID: 3, MetaInfo: "Record 3!"},
			wantState: getStateMap(binaryDataState),
			wantError: ErrEmptyData,
		},
		{
			name: "Test 5. Update text data, change metaInfo with non empty content",
			td:   storage.BinaryData{ID: 1, BinaryData: []byte("Raiden Shogun"), MetaInfo: "Record updated 1!"},
			wantState: map[int]storage.BinaryData{
				1: {ID: 1, BinaryData: []byte("Raiden Shogun"), MetaInfo: "Record updated 1!"},
				3: {ID: 3, BinaryData: []byte("Ayato")},
			},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBinaryData{BinaryData: getStateMap(binaryDataState), LastUsedID: lastUsedID}

			gotError := rep.UpdateBinaryData(ctx, tt.td)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BinaryData)
			assert.Equal(t, rep.LastUsedID, lastUsedID)
		})
	}
}

func TestMockBinaryData_DeleteBinaryData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		td        storage.BinaryData
		wantState map[int]storage.BinaryData
		wantError error
	}{
		{
			name: "Test 1. Delete text data, text data exist",
			td:   storage.BinaryData{ID: 1},
			wantState: map[int]storage.BinaryData{
				3: {ID: 3, BinaryData: []byte("Ayato")},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Delete text data, text data not exist",
			td:        storage.BinaryData{},
			wantState: getStateMap(binaryDataState),
			wantError: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBinaryData{BinaryData: getStateMap(binaryDataState), LastUsedID: lastUsedID}

			gotError := rep.DeleteBinaryData(ctx, tt.td)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BinaryData)
			assert.Equal(t, rep.LastUsedID, lastUsedID)
		})
	}
}
