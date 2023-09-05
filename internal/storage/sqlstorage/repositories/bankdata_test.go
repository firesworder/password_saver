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

func clearBankDataTable(t *testing.T, db *sql.DB) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM bankdata")
	require.NoError(t, err)
}

func TestBankDataCommon(t *testing.T) {
	var err error
	ctx := context.Background()
	sqlS, err := sqlstorage.NewStorage(storage.DevDSN)
	conn := sqlS.Connection
	require.NoError(t, err)

	defer clearUserTable(t, conn)
	defer clearBankDataTable(t, conn)

	uRep := User{Conn: conn}
	user1, err := uRep.CreateUser(ctx, storage.User{Login: "Ayaka", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	user2, err := uRep.CreateUser(ctx, storage.User{Login: "Ayato", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	uid1, uid2 := user1.ID, user2.ID

	bdRep := BankData{Conn: conn}
	bid1, err := bdRep.AddBankData(ctx, storage.BankData{
		CardNumber: "1122334455667788", CardExpire: "11/23", CVV: "852", MetaInfo: "Ayaka note", UserID: uid1,
	})
	require.NoError(t, err)
	bid2, err := bdRep.AddBankData(ctx, storage.BankData{
		CardNumber: "8855664411228800", CardExpire: "05/25", CVV: "146", MetaInfo: "Ayato note", UserID: uid2,
	})
	require.NoError(t, err)

	err = bdRep.UpdateBankData(ctx, storage.BankData{ID: bid1, CardNumber: "5544556655225588", CardExpire: "11/24",
		CVV: "789", MetaInfo: "Ayaka updated note", UserID: user1.ID,
	})
	require.NoError(t, err)

	err = bdRep.DeleteBankData(ctx, storage.BankData{ID: bid2})
	require.NoError(t, err)

	// todo: доделать GetAllRecords после фикса привязки по UserID
	tdList1, err := bdRep.GetAllRecords(ctx)
	assert.Equal(t, []storage.BankData{
		{ID: bid1, CardNumber: "5544556655225588", CardExpire: "11/24", CVV: "789", MetaInfo: "Ayaka updated note",
			UserID: uid1},
	}, tdList1)
}
