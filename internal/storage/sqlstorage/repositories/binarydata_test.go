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

func clearBinaryDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM binarydata")
	require.NoError(t, err)
}

func TestBinaryDataCommon(t *testing.T) {
	var err error
	ctx := context.Background()
	sqlS, err := sqlstorage.NewStorage(storage.DevDSN)
	conn := sqlS.Connection
	require.NoError(t, err)

	defer clearUserTable(t, conn)
	defer clearBinaryDataTable(t, conn)

	uRep := User{Conn: conn}
	user1, err := uRep.CreateUser(ctx, storage.User{Login: "Ayaka", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	user2, err := uRep.CreateUser(ctx, storage.User{Login: "Ayato", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	uid1, uid2 := user1.ID, user2.ID

	bdRep := BinaryData{Conn: conn}
	bid1, err := bdRep.AddBinaryData(ctx,
		storage.BinaryData{BinaryData: []byte("Ayaka note"), MetaInfo: "td1", UserID: uid1})
	require.NoError(t, err)
	bid2, err := bdRep.AddBinaryData(ctx,
		storage.BinaryData{BinaryData: []byte("Ayato note"), MetaInfo: "td2", UserID: uid2})
	require.NoError(t, err)

	err = bdRep.UpdateBinaryData(ctx,
		storage.BinaryData{ID: bid1, BinaryData: []byte("Ayaka updated!"), MetaInfo: "updtd1"})
	require.NoError(t, err)

	err = bdRep.DeleteBinaryData(ctx, storage.BinaryData{ID: bid2})
	require.NoError(t, err)

	// todo: доделать GetAllRecords после фикса привязки по UserID
	tdList1, err := bdRep.GetAllRecords(ctx)
	assert.Equal(t, []storage.BinaryData{
		{ID: bid1, BinaryData: []byte("Ayaka updated!"), MetaInfo: "updtd1", UserID: uid1},
	}, tdList1)
}
