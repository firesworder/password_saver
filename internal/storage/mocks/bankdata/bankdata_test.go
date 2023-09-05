package bankdata

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var bankDataState = map[int]storage.BankData{
	1: {ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "123"},
	3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
}
var lastUsedID = 3

func getStateMap(src map[int]storage.BankData) map[int]storage.BankData {
	r := map[int]storage.BankData{}
	for key, value := range src {
		r[key] = value
	}
	return r
}

func Test_checkBankData(t *testing.T) {
	tests := []struct {
		name      string
		bd        storage.BankData
		wantError error
	}{
		{
			name:      "Test 1. Correct bank data",
			bd:        storage.BankData{CardNumber: "7788 9900 1122 3344", CardExpire: "09/23", CVV: "123"},
			wantError: nil,
		},
		{
			name:      "Test 2. All fields are empty",
			bd:        storage.BankData{},
			wantError: ErrDataInvalid,
		},
		{
			name:      "Test 3. One bank data(except ID and MetaInfo) of fields is empty",
			bd:        storage.BankData{CardNumber: "7788 9900 1122 3344", CardExpire: "09/23", CVV: ""},
			wantError: ErrDataInvalid,
		},
		{
			name:      "Test 4. Incorrect card number",
			bd:        storage.BankData{CardNumber: "7788 9900 1122 3344 5566", CardExpire: "09/23", CVV: "123"},
			wantError: ErrDataInvalid,
		},
		{
			name:      "Test 5. Incorrect card expire",
			bd:        storage.BankData{CardNumber: "7788 9900 1122 3344", CardExpire: "d12/23", CVV: "123"},
			wantError: ErrDataInvalid,
		},
		{
			name:      "Test 6. Incorrect CVV",
			bd:        storage.BankData{CardNumber: "7788 9900 1122 3344", CardExpire: "d12/23", CVV: "12355"},
			wantError: ErrDataInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotError := checkBankData(tt.bd)
			assert.ErrorIs(t, gotError, tt.wantError)
		})
	}
}

func TestMockBankData_AddBankData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		bd             storage.BankData
		wantState      map[int]storage.BankData
		wantLastUsedID int
		wantError      error
	}{
		{
			name: "Test 1. Correct bank data",
			bd:   storage.BankData{CardNumber: "9911 8822 7733 6644", CardExpire: "02/22", CVV: "808"},
			wantState: map[int]storage.BankData{
				1: {ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "123"},
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
				4: {ID: 4, CardNumber: "9911 8822 7733 6644", CardExpire: "02/22", CVV: "808"},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name: "Test 2. Correct bank data, with metaInfo",
			bd: storage.BankData{
				CardNumber: "9911 8822 7733 6644", CardExpire: "02/22", CVV: "808", MetaInfo: "bank data meta info",
			},
			wantState: map[int]storage.BankData{
				1: {ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "123"},
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
				4: {
					ID: 4, CardNumber: "9911 8822 7733 6644", CardExpire: "02/22", CVV: "808",
					MetaInfo: "bank data meta info",
				},
			},
			wantLastUsedID: 4,
			wantError:      nil,
		},
		{
			name: "Test 3. Bank data with error(s) in fields",
			bd:   storage.BankData{CardNumber: "9911 8822 7733", CardExpire: "02///22", CVV: "808"},
			wantState: map[int]storage.BankData{
				1: {ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "123"},
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
			},
			wantLastUsedID: 3,
			wantError:      ErrDataInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBankData{BankData: getStateMap(bankDataState), LastUsedID: lastUsedID}

			id, gotError := rep.AddBankData(ctx, tt.bd)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BankData)
			if gotError == nil {
				assert.Equal(t, tt.wantLastUsedID, id)
			}
			assert.Equal(t, tt.wantLastUsedID, rep.LastUsedID)
		})
	}
}

func TestMockBankData_UpdateBankData(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		bd        storage.BankData
		wantState map[int]storage.BankData
		wantError error
	}{
		{
			name: "Test 1. Correct update bank data(change CVV)",
			bd:   storage.BankData{ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "000"},
			wantState: map[int]storage.BankData{
				1: {ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "000"},
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Update bank data with error, ID unknown",
			bd:        storage.BankData{ID: 10, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "000"},
			wantState: getStateMap(bankDataState),
			wantError: ErrNotFound,
		},
		{
			name:      "Test 3. Update bank data with error, error(s) in fields(card expire invalid)",
			bd:        storage.BankData{ID: 10, CardNumber: "0099 8877 6655 4433", CardExpire: "d09///23", CVV: "000"},
			wantState: getStateMap(bankDataState),
			wantError: ErrDataInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBankData{BankData: getStateMap(bankDataState), LastUsedID: lastUsedID}

			gotError := rep.UpdateBankData(ctx, tt.bd)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BankData)
		})
	}
}

func TestMockBankData_DeleteBankData(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name      string
		bd        storage.BankData
		wantState map[int]storage.BankData
		wantError error
	}{
		{
			name: "Test 1. Correct delete bank data",
			bd:   storage.BankData{ID: 1},
			wantState: map[int]storage.BankData{
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
			},
			wantError: nil,
		},
		{
			name:      "Test 2. Delete bank data with error, ID unknown",
			bd:        storage.BankData{ID: 10, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "000"},
			wantState: getStateMap(bankDataState),
			wantError: ErrNotFound,
		},
		{
			name: "Test 3. Correct delete, with error in excessive fields",
			bd:   storage.BankData{ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "d09///23", CVV: "000"},
			wantState: map[int]storage.BankData{
				3: {ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
			},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBankData{BankData: getStateMap(bankDataState), LastUsedID: lastUsedID}

			gotError := rep.DeleteBankData(ctx, tt.bd)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantState, rep.BankData)
		})
	}
}

func TestMockBankData_GetAllRecords(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		state        map[int]storage.BankData
		wantRecSlice []storage.BankData
		wantError    error
	}{
		{
			name:         "Test 1. Empty state",
			state:        map[int]storage.BankData{},
			wantRecSlice: []storage.BankData{},
			wantError:    nil,
		},
		{
			name:  "Test 2. Filled state",
			state: getStateMap(bankDataState),
			wantRecSlice: []storage.BankData{
				{ID: 1, CardNumber: "0099 8877 6655 4433", CardExpire: "09/23", CVV: "123"},
				{ID: 3, CardNumber: "1122 3344 5566 7788", CardExpire: "06/24", CVV: "987"},
			},
			wantError: nil,
		},
		{
			name:         "Test 3. nil state",
			state:        nil,
			wantRecSlice: []storage.BankData{},
			wantError:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockBankData{BankData: getStateMap(tt.state), LastUsedID: lastUsedID}

			gotRecords, gotErr := rep.GetAllRecords(ctx)
			sort.Slice(gotRecords, func(i, j int) bool {
				return gotRecords[i].ID < gotRecords[j].ID
			})
			assert.EqualValues(t, tt.wantRecSlice, gotRecords)
			assert.Equal(t, tt.wantError, gotErr)
		})
	}
}
