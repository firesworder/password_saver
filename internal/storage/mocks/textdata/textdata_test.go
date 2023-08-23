package textdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"password_saver/internal/storage"
	"testing"
)

var textDataState = map[int]storage.TextData{
	1: {ID: 1, TextData: "Hello world!"},
	3: {ID: 3, TextData: "Ayayaka"},
}
var lastUsedID = 3

func getStateMap(src map[int]storage.TextData) map[int]storage.TextData {
	r := map[int]storage.TextData{}
	for key, value := range src {
		r[key] = value
	}
	return r
}

func TestMockTextData_AddTextData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		td             storage.TextData
		wantState      map[int]storage.TextData
		wantLastUsedID int
		wantError      error
	}{
		{
			name: "Test 1. Add text data, without ID",
			td:   storage.TextData{TextData: "Ahoy! Me Hearties"},
			wantState: map[int]storage.TextData{
				1: {ID: 1, TextData: "Hello world!"},
				3: {ID: 3, TextData: "Ayayaka"},
				4: {ID: 4, TextData: "Ahoy! Me Hearties"},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name: "Test 2. Add text data, with ID",
			td:   storage.TextData{ID: 10, TextData: "Ahoy! Me Hearties"},
			wantState: map[int]storage.TextData{
				1: {ID: 1, TextData: "Hello world!"},
				3: {ID: 3, TextData: "Ayayaka"},
				4: {ID: 4, TextData: "Ahoy! Me Hearties"},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name:           "Test 3. Add empty text data",
			td:             storage.TextData{TextData: ""},
			wantState:      getStateMap(textDataState),
			wantLastUsedID: lastUsedID,
			wantError:      ErrEmptyData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockTextData{textDataMap: getStateMap(textDataState), lastUsedID: lastUsedID}

			gotError := rep.AddTextData(ctx, tt.td)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.textDataMap)
			assert.Equal(t, tt.wantLastUsedID, rep.lastUsedID)
		})
	}
}

func TestMockTextData_UpdateTextData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		td        storage.TextData
		wantState map[int]storage.TextData
		wantError error
	}{
		{
			name: "Test 1. Update text data",
			td:   storage.TextData{ID: 1, TextData: "Neeko is the best decision!"},
			wantState: map[int]storage.TextData{
				1: {ID: 1, TextData: "Neeko is the best decision!"},
				3: {ID: 3, TextData: "Ayayaka"},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Update text data, ID not set(id=0)",
			td:        storage.TextData{TextData: "I will...Neeko!"},
			wantState: getStateMap(textDataState),
			wantError: ErrNotFound,
		},
		{
			name:      "Test 3. Update text data, empty data",
			td:        storage.TextData{ID: 3, TextData: ""},
			wantState: getStateMap(textDataState),
			wantError: ErrEmptyData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockTextData{textDataMap: getStateMap(textDataState), lastUsedID: lastUsedID}

			gotError := rep.UpdateTextData(ctx, tt.td)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.textDataMap)
			assert.Equal(t, rep.lastUsedID, lastUsedID)
		})
	}
}

func TestMockTextData_DeleteTextData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		td        storage.TextData
		wantState map[int]storage.TextData
		wantError error
	}{
		{
			name: "Test 1. Delete text data, text data exist",
			td:   storage.TextData{ID: 1},
			wantState: map[int]storage.TextData{
				3: {ID: 3, TextData: "Ayayaka"},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Delete text data, text data not exist",
			td:        storage.TextData{},
			wantState: getStateMap(textDataState),
			wantError: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockTextData{textDataMap: getStateMap(textDataState), lastUsedID: lastUsedID}

			gotError := rep.DeleteTextData(ctx, tt.td)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.textDataMap)
			assert.Equal(t, rep.lastUsedID, lastUsedID)
		})
	}
}
