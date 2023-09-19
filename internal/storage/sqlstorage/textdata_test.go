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

func TextDataRep(t *testing.T) *TextData {
	s := devStorage(t)
	enc, err := crypt.NewEncoder("tests/cert_test.pem")
	require.NoError(t, err)
	dec, err := crypt.NewDecoder("tests/privKey_test.pem")
	require.NoError(t, err)
	tdRep := TextData{Conn: s.Connection, Encoder: enc, Decoder: dec}
	return &tdRep
}

func TestTextData(t *testing.T) {
	ctx := context.Background()

	tdRep := TextDataRep(t)
	// закрываю соединение с дб
	defer tdRep.Conn.Close()

	// очищаю таблицы перед добавлением новых тестовых данных и по итогам прогона тестов
	clearTables(t, tdRep.Conn)
	defer clearTables(t, tdRep.Conn)

	// подготовка тестовых данных
	var uID int64
	var err error
	err = tdRep.Conn.QueryRowContext(ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) returning id", "demo", "demo").Scan(&uID)
	require.NoError(t, err)
	testUser := storage.User{ID: int(uID), Login: "demo", HashedPassword: "demo"}

	// добавление записи
	id1, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "text data 1", MetaInfo: "meta info 1"}, &testUser)
	require.NoError(t, err)
	id2, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "text data 2", MetaInfo: "meta info 2"}, &testUser)
	require.NoError(t, err)
	id3, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "text data 3", MetaInfo: "meta info 3"}, &testUser)
	require.NoError(t, err)

	// обновление записи
	err = tdRep.UpdateTextData(ctx,
		storage.TextData{ID: id2, TextData: "upd text data", MetaInfo: "upd meta info"}, &testUser)
	require.NoError(t, err)
	err = tdRep.UpdateTextData(ctx, storage.TextData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// удаление записи
	err = tdRep.DeleteTextData(ctx, storage.TextData{ID: id3}, &testUser)
	require.NoError(t, err)
	err = tdRep.DeleteTextData(ctx, storage.TextData{ID: 0}, &testUser)
	require.ErrorIs(t, err, storage.ErrElementNotFound)

	// вывод записей
	gotResult, err := tdRep.GetAllRecords(ctx, &testUser)
	require.NoError(t, err)
	sort.Slice(gotResult, func(i, j int) bool {
		return gotResult[i].ID < gotResult[j].ID
	})
	wantResult := []storage.TextData{
		{ID: id1, TextData: "text data 1", MetaInfo: "meta info 1", UserID: int(uID)},
		{ID: id2, TextData: "upd text data", MetaInfo: "upd meta info", UserID: int(uID)},
	}
	assert.Equal(t, wantResult, gotResult)
}
