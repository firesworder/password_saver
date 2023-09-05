package repositories

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTextDataCommon(t *testing.T) {
	var err error
	ctx := context.Background()
	conn := getConnection(t)

	defer clearUserTable(t, conn)
	defer clearTextDataTable(t, conn)

	uRep := User{Conn: conn}
	user1, err := uRep.CreateUser(ctx, storage.User{Login: "Ayaka", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	user2, err := uRep.CreateUser(ctx, storage.User{Login: "Ayato", HashedPassword: "Kamisato"})
	require.NoError(t, err)
	uid1, uid2 := user1.ID, user2.ID

	tdRep := TextData{Conn: conn}
	tid1, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "Ayaka note", MetaInfo: "td1", UserID: uid1})
	require.NoError(t, err)
	tid2, err := tdRep.AddTextData(ctx, storage.TextData{TextData: "Ayato note", MetaInfo: "td2", UserID: uid2})
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
