package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func clearTextDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM textdata")
	require.NoError(t, err)
}

func TestCommon(t *testing.T) {
	var err error
	ctx := context.Background()
	sqlS, err := sqlstorage.NewStorage(storage.DevDSN)
	conn := sqlS.Connection
	require.NoError(t, err)

	defer clearUserTable(t, conn)
	defer clearTextDataTable(t, conn)

	uRep := User{conn: conn}
	user1, err := uRep.CreateUser(ctx, storage.User{Login: "Ayaka", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	user2, err := uRep.CreateUser(ctx, storage.User{Login: "Ayato", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	uid1, _ := user1.ID, user2.ID

	tdRep := TextData{conn: conn}
	td1 := storage.TextData{TextData: "Ayaka note", MetaInfo: "td1", UserID: uid1}
	tid1, err := tdRep.AddTextData(ctx, td1)
	require.NoError(t, err)
	tid2, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "Ayato note", MetaInfo: "td2", UserID: uid1})
	require.NoError(t, err)

	err = tdRep.UpdateTextData(ctx, storage.TextData{ID: tid1, TextData: "Ayaka updated!", MetaInfo: "updtd1"})
	require.NoError(t, err)

	err = tdRep.DeleteTextData(ctx, storage.TextData{ID: tid2})
	require.NoError(t, err)

	// todo: доделать GetAllRecords после фикса привязки по UserID
	tdList1, err := tdRep.GetAllRecords(ctx)
	assert.Equal(t, []storage.TextData{
		{ID: tid1, TextData: "Ayaka updated!", MetaInfo: "updtd1", UserID: uid1},
	}, tdList1)
}
