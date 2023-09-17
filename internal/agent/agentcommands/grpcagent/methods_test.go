package grpcagent

import (
	"context"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func startTestServer(t *testing.T) (pb.PasswordSaverClient, func()) {
	serverStarted := make(chan struct{})
	// определяем порт для сервера
	listen := bufconn.Listen(1024 * 1024)

	// инциал. сервис
	service := &mocks.GRPCServer{}
	// создаем пустой grpc сервер, без опций
	server := grpc.NewServer()
	// регистрируем сервис на сервере
	pb.RegisterPasswordSaverServer(server, service)

	// запуск сервера в горутине
	go func() {
		serverStarted <- struct{}{}
		// запускаем grpc сервер на выделенном порту 'listen'
		if err := server.Serve(listen); err != nil {
			require.NoError(t, err)
		}
	}()
	<-serverStarted

	bufConnDialer := func(ctx context.Context, s string) (net.Conn, error) {
		return listen.Dial()
	}

	// создание соединения к запущенному серверу
	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(bufConnDialer), grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	client := pb.NewPasswordSaverClient(conn)

	closer := func() {
		var err error
		// закрываю grpc сервер
		server.GracefulStop()

		err = conn.Close()
		require.NoError(t, err)
	}
	return client, closer
}

func TestTextMethods(t *testing.T) {
	var err error
	client, closer := startTestServer(t)
	defer closer()
	ga := GRPCAgent{grpcClient: client}

	// create_text_record
	id, err := ga.CreateTextDataRecord(storage.TextData{ID: 100})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	id, err = ga.CreateTextDataRecord(storage.TextData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 0, id)

	// update_text_record
	err = ga.UpdateTextDataRecord(storage.TextData{ID: 100, TextData: "td1"})
	assert.NoError(t, err)

	err = ga.UpdateTextDataRecord(storage.TextData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)

	// delete_text_record
	err = ga.DeleteTextDataRecord(storage.TextData{ID: 50, TextData: "td1"})
	assert.NoError(t, err)

	err = ga.DeleteTextDataRecord(storage.TextData{ID: 100})
	assert.NotEmpty(t, err)
}

func TestBankMethods(t *testing.T) {
	var err error
	client, closer := startTestServer(t)
	defer closer()
	ga := GRPCAgent{grpcClient: client}

	// create_bank_record
	id, err := ga.CreateBankDataRecord(storage.BankData{ID: 100})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	id, err = ga.CreateBankDataRecord(storage.BankData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 0, id)

	// update_bank_record
	err = ga.UpdateBankDataRecord(storage.BankData{ID: 100, CVV: "342"})
	assert.NoError(t, err)

	err = ga.UpdateBankDataRecord(storage.BankData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)

	// delete_bank_record
	err = ga.DeleteBankDataRecord(storage.BankData{ID: 50, CVV: "463"})
	assert.NoError(t, err)

	err = ga.DeleteBankDataRecord(storage.BankData{ID: 100})
	assert.NotEmpty(t, err)

}

func TestBinaryMethods(t *testing.T) {
	var err error
	client, closer := startTestServer(t)
	defer closer()
	ga := GRPCAgent{grpcClient: client}

	// create_binary_record
	id, err := ga.CreateBinaryDataRecord(storage.BinaryData{ID: 100})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	id, err = ga.CreateBinaryDataRecord(storage.BinaryData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 0, id)

	// update_binary_record
	err = ga.UpdateBinaryDataRecord(storage.BinaryData{ID: 100, MetaInfo: "mi2"})
	assert.NoError(t, err)

	err = ga.UpdateBinaryDataRecord(storage.BinaryData{ID: 100, MetaInfo: "skip"})
	assert.NotEmpty(t, err)

	// delete_binary_record
	err = ga.DeleteBinaryDataRecord(storage.BinaryData{ID: 50, MetaInfo: "mi2"})
	assert.NoError(t, err)

	err = ga.DeleteBinaryDataRecord(storage.BinaryData{ID: 100})
	assert.NotEmpty(t, err)

}

func TestOtherMethods(t *testing.T) {
	var err error
	client, closer := startTestServer(t)
	defer closer()
	ga := GRPCAgent{grpcClient: client}

	// register_user
	err = ga.RegisterUser("demoLog", "demoPass")
	assert.NoError(t, err)
	assert.Equal(t, "demo_token", ga.userToken)
	ga.userToken = ""

	err = ga.RegisterUser("admin", "admin")
	assert.NotEmpty(t, err)
	assert.Equal(t, "", ga.userToken)

	// login_user
	err = ga.LoginUser("demoLog", "demoPass")
	assert.NoError(t, err)
	assert.Equal(t, "demo_token", ga.userToken)
	ga.userToken = ""

	err = ga.LoginUser("admin", "admin")
	assert.NotEmpty(t, err)
	assert.Equal(t, "", ga.userToken)

	// get_all_records
	stg, err := ga.ShowAllRecords()
	assert.NoError(t, err)
	assert.NotEmpty(t, stg.TextDataList)
}
