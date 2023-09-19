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

func BinaryDataRep(t *testing.T) *BinaryData {
	s := devStorage(t)
	enc, err := crypt.NewEncoder("tests/cert_test.pem")
	require.NoError(t, err)
	dec, err := crypt.NewDecoder("tests/privKey_test.pem")
	require.NoError(t, err)
	bdRep := BinaryData{Conn: s.Connection, Encoder: enc, Decoder: dec}
	return &bdRep
}

func TestBinaryData(t *testing.T) {
	ctx := context.Background()

	bdRep := BinaryDataRep(t)
	// закрываю соединение с дб
	defer bdRep.Conn.Close()

	// очищаю таблицы перед добавлением новых тестовых данных и по итогам прогона тестов
	clearTables(t, bdRep.Conn)
	defer clearTables(t, bdRep.Conn)

	// подготовка тестовых данных
	var uID int64
	var err error
	err = bdRep.Conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo", "demo").Scan(&uID)
	require.NoError(t, err)
	testUser := storage.User{ID: int(uID), Login: "demo", HashedPassword: "demo"}

	// добавление записи
	id1, err := bdRep.AddBinaryData(ctx,
		storage.BinaryData{BinaryData: []byte("binary data 1"), MetaInfo: "meta info 1"}, &testUser)
	require.NoError(t, err)
	id2, err := bdRep.AddBinaryData(ctx,
		storage.BinaryData{BinaryData: []byte("binary data 2"), MetaInfo: "meta info 2"}, &testUser)
	require.NoError(t, err)
	id3, err := bdRep.AddBinaryData(ctx,
		storage.BinaryData{BinaryData: []byte("binary data 3"), MetaInfo: "meta info 3"}, &testUser)
	require.NoError(t, err)

	// обновление записи
	err = bdRep.UpdateBinaryData(ctx,
		storage.BinaryData{ID: id2, BinaryData: []byte("upd binary data"), MetaInfo: "upd meta info"}, &testUser)
	require.NoError(t, err)
	err = bdRep.UpdateBinaryData(ctx, storage.BinaryData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// удаление записи
	err = bdRep.DeleteBinaryData(ctx, storage.BinaryData{ID: id3}, &testUser)
	require.NoError(t, err)
	err = bdRep.DeleteBinaryData(ctx, storage.BinaryData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// вывод записей
	gotResult, err := bdRep.GetAllRecords(ctx, &testUser)
	require.NoError(t, err)
	sort.Slice(gotResult, func(i, j int) bool {
		return gotResult[i].ID < gotResult[j].ID
	})
	wantResult := []storage.BinaryData{
		{ID: id1, BinaryData: []byte("binary data 1"), MetaInfo: "meta info 1", UserID: int(uID)},
		{ID: id2, BinaryData: []byte("upd binary data"), MetaInfo: "upd meta info", UserID: int(uID)},
	}
	assert.Equal(t, wantResult, gotResult)
}
