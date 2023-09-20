package sqlstorage

import (
	"context"
	"github.com/firesworder/password_saver/internal/crypt"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func BankDataRep(t *testing.T) *BankData {
	s := devStorage(t)
	enc, err := crypt.NewEncoder("tests/cert_test.pem")
	require.NoError(t, err)
	dec, err := crypt.NewDecoder("tests/privKey_test.pem")
	require.NoError(t, err)
	bdRep := BankData{Conn: s.Connection, Encoder: enc, Decoder: dec}
	return &bdRep
}

func TestBankData(t *testing.T) {
	ctx := context.Background()

	bdRep := BankDataRep(t)
	// закрываю соединение с дб
	defer bdRep.Conn.Close()

	// очищаю таблицы перед добавлением новых тестовых данных и по итогам прогона тестов
	clearTables(t, bdRep.Conn)
	defer clearTables(t, bdRep.Conn)

	// подготовка тестовых данных
	var uID, uID2 int64
	var err error
	err = bdRep.Conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo", "demo").Scan(&uID)
	require.NoError(t, err)
	testUser := storage.User{ID: int(uID), Login: "demo", HashedPassword: "demo"}
	// testUser2
	err = bdRep.Conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo2", "demo2").Scan(&uID2)
	require.NoError(t, err)
	testUser2 := storage.User{ID: int(uID2), Login: "demo2", HashedPassword: "demo2"}

	// добавление записи
	id1, err := bdRep.AddBankData(ctx,
		storage.BankData{CardNumber: "0011223344556677", CardExpire: "09/25", CVV: "333"}, &testUser)
	require.NoError(t, err)
	id2, err := bdRep.AddBankData(ctx,
		storage.BankData{CardNumber: "1133553366331133", CardExpire: "09/26", CVV: "444"}, &testUser2)
	require.NoError(t, err)
	id3, err := bdRep.AddBankData(ctx,
		storage.BankData{CardNumber: "5500110022003300", CardExpire: "03/25", CVV: "555"}, &testUser)
	require.NoError(t, err)
	id4, err := bdRep.AddBankData(ctx,
		storage.BankData{CardNumber: "9911882277336644", CardExpire: "06/24", CVV: "462"}, &testUser)
	require.NoError(t, err)

	// обновление записи
	err = bdRep.UpdateBankData(ctx,
		storage.BankData{ID: id2, CardNumber: "3333222211114444", CardExpire: "03/23", CVV: "123"}, &testUser2)
	require.NoError(t, err)
	err = bdRep.UpdateBankData(ctx, storage.BankData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// удаление записи
	err = bdRep.DeleteBankData(ctx, storage.BankData{ID: id3}, &testUser)
	require.NoError(t, err)
	err = bdRep.DeleteBankData(ctx, storage.BankData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// вывод записей
	// testUser
	gotResult, err := bdRep.GetAllRecords(ctx, &testUser)
	require.NoError(t, err)
	sort.Slice(gotResult, func(i, j int) bool {
		return gotResult[i].ID < gotResult[j].ID
	})
	wantResult := []storage.BankData{
		{ID: id1, CardNumber: "0011223344556677", CardExpire: "09/25", CVV: "333", UserID: int(uID)},
		{ID: id4, CardNumber: "9911882277336644", CardExpire: "06/24", CVV: "462", UserID: int(uID)},
	}
	sort.Slice(gotResult, func(i, j int) bool {
		return gotResult[i].ID < gotResult[j].ID
	})
	assert.Equal(t, wantResult, gotResult)
	//testUser2
	gotResult, err = bdRep.GetAllRecords(ctx, &testUser2)
	require.NoError(t, err)
	wantResult = []storage.BankData{
		{ID: id2, CardNumber: "3333222211114444", CardExpire: "03/23", CVV: "123", UserID: int(uID2)},
	}
	assert.Equal(t, wantResult, gotResult)
}
