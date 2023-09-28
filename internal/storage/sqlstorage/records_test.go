package sqlstorage

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestRecordRepository(t *testing.T) {
	ctx := context.Background()
	s := devStorage(t)
	conn := s.Connection
	rr := RecordRepository{conn: conn}

	defer s.Connection.Close()

	// очистка таблицы до и после изменений
	clearTables(t, conn)
	defer clearTables(t, conn)

	// добавление пользователей
	var uid1, uid2 int
	err := conn.QueryRowContext(ctx, "INSERT INTO users(login, password) VALUES($1, $2) RETURNING id",
		"user1", "upass1").Scan(&uid1)
	require.NoError(t, err)
	err = conn.QueryRowContext(ctx, "INSERT INTO users(login, password) VALUES($1, $2) RETURNING id",
		"user2", "upass2").Scan(&uid2)
	require.NoError(t, err)

	// добавление записей
	r1 := storage.Record{RecordType: "text", Content: []byte("text content"), MetaInfo: "MI1"}
	rid1, err := rr.AddRecord(ctx, r1, uid1)
	require.NoError(t, err)
	rid2, err := rr.AddRecord(ctx,
		storage.Record{RecordType: "bank", Content: []byte("bank data content"), MetaInfo: "MI2"}, uid2)
	require.NoError(t, err)
	rid3, err := rr.AddRecord(ctx,
		storage.Record{RecordType: "binary", Content: []byte("binary content"), MetaInfo: "MI3"}, uid1)
	require.NoError(t, err)
	r4 := storage.Record{RecordType: "binary", Content: []byte("binary content"), MetaInfo: "MI4"}
	rid4, err := rr.AddRecord(ctx, r4, uid1)
	require.NoError(t, err)

	// обновление записей
	updR2 := storage.Record{ID: rid2, RecordType: "bank", Content: []byte("upd bank data content"), MetaInfo: "upd MI2"}
	err = rr.UpdateRecord(ctx, updR2, uid2)
	require.NoError(t, err)
	// ошибка обновления - неправильный пользователь у записи(должен быть uid2)
	err = rr.UpdateRecord(ctx, updR2, uid1)
	assert.ErrorIs(t, err, storage.ErrElementNotFound)
	// ошибка обновления - неизвестный id у записи
	updR2Err := updR2
	updR2Err.ID += 100
	err = rr.UpdateRecord(ctx, updR2Err, uid2)
	assert.ErrorIs(t, err, storage.ErrElementNotFound)

	// удаление записи
	err = rr.DeleteRecord(ctx, storage.Record{ID: rid3}, uid1)
	require.NoError(t, err)
	// ошибка обновления - неправильный пользователь у записи(должен быть uid1)
	err = rr.DeleteRecord(ctx, storage.Record{ID: rid1}, uid2)
	assert.ErrorIs(t, err, storage.ErrElementNotFound)
	// ошибка обновления - неизвестный id у записи
	err = rr.DeleteRecord(ctx, storage.Record{ID: rid3 + 100}, uid1)
	assert.ErrorIs(t, err, storage.ErrElementNotFound)

	// получение всех записей
	// пользователь uid1
	recordSlice1, err := rr.GetAll(ctx, uid1)
	require.NoError(t, err)
	sort.Slice(recordSlice1, func(i, j int) bool {
		return recordSlice1[i].ID < recordSlice1[j].ID
	})
	r1.ID, r4.ID = rid1, rid4
	wantResult1 := []storage.Record{r1, r4}
	assert.Equal(t, wantResult1, recordSlice1)

	// пользователь uid2
	recordSlice2, err := rr.GetAll(ctx, uid2)
	require.NoError(t, err)
	assert.Equal(t, []storage.Record{updR2}, recordSlice2)
}
