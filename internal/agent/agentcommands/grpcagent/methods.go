package grpcagent

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc/metadata"
)

const ctxTokenParam = "userToken"

// RegisterUser регистрирует пользователя.
func (a *GRPCAgent) RegisterUser(login, password string) error {
	req := pb.RegisterUserRequest{Login: login, Password: password}

	resp, err := a.grpcClient.RegisterUser(context.Background(), &req)
	if err != nil {
		return err
	}
	a.userToken = resp.Token
	return nil
}

// LoginUser авторизует пользователя.
func (a *GRPCAgent) LoginUser(login, password string) error {
	req := pb.LoginUserRequest{Login: login, Password: password}

	resp, err := a.grpcClient.LoginUser(context.Background(), &req)
	if err != nil {
		return err
	}
	a.userToken = resp.Token
	return nil
}

// CreateTextDataRecord создает запись текстовых данных.
func (a *GRPCAgent) CreateTextDataRecord(input storage.TextData) (int, error) {
	req := pb.AddTextDataRequest{TextData: storage.TextDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	resp, err := a.grpcClient.AddTextDataRecord(ctx, &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

// CreateBankDataRecord создает запись банковских данных.
func (a *GRPCAgent) CreateBankDataRecord(input storage.BankData) (int, error) {
	req := pb.AddBankDataRequest{BankData: storage.BankDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	resp, err := a.grpcClient.AddBankDataRecord(ctx, &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

// CreateBinaryDataRecord создает запись бинарных данных.
func (a *GRPCAgent) CreateBinaryDataRecord(input storage.BinaryData) (int, error) {
	req := pb.AddBinaryDataRequest{BinaryData: storage.BinaryDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	resp, err := a.grpcClient.AddBinaryDataRecord(ctx, &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

// UpdateTextDataRecord обновляет запись тестовых данных.
func (a *GRPCAgent) UpdateTextDataRecord(input storage.TextData) error {
	req := pb.UpdateTextDataRequest{TextData: storage.TextDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.UpdateTextDataRecord(ctx, &req)
	return err
}

// UpdateBankDataRecord обновляет запись банковских данных.
func (a *GRPCAgent) UpdateBankDataRecord(input storage.BankData) error {
	req := pb.UpdateBankDataRequest{BankData: storage.BankDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.UpdateBankDataRecord(ctx, &req)
	return err
}

// UpdateBinaryDataRecord обновляет запись бинарных данных.
func (a *GRPCAgent) UpdateBinaryDataRecord(input storage.BinaryData) error {
	req := pb.UpdateBinaryDataRequest{BinaryData: storage.BinaryDataToGRPC(input)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.UpdateBinaryDataRecord(ctx, &req)
	return err
}

// DeleteTextDataRecord удаляет запись текстовых данных.
func (a *GRPCAgent) DeleteTextDataRecord(input storage.TextData) error {
	req := pb.DeleteTextDataRequest{Id: int64(input.ID)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.DeleteTextDataRecord(ctx, &req)
	return err
}

// DeleteBankDataRecord удаляет запись банковских данных.
func (a *GRPCAgent) DeleteBankDataRecord(input storage.BankData) error {
	req := pb.DeleteBankDataRequest{Id: int64(input.ID)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.DeleteBankDataRecord(ctx, &req)
	return err
}

// DeleteBinaryDataRecord удаляет запись бинарных данных.
func (a *GRPCAgent) DeleteBinaryDataRecord(input storage.BinaryData) error {
	req := pb.DeleteBinaryDataRequest{Id: int64(input.ID)}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	_, err := a.grpcClient.DeleteBinaryDataRecord(ctx, &req)
	return err
}

// ShowAllRecords получает список записей с сервера.
func (a *GRPCAgent) ShowAllRecords() (*storage.RecordsList, error) {
	req := pb.GetAllRecordsRequest{}

	ctx := metadata.NewOutgoingContext(
		context.Background(), metadata.New(map[string]string{ctxTokenParam: a.userToken}))
	resp, err := a.grpcClient.GetAllRecords(ctx, &req)
	if err != nil {
		return nil, err
	}

	result := storage.RecordsList{
		TextDataList:   []storage.TextData{},
		BankDataList:   []storage.BankData{},
		BinaryDataList: []storage.BinaryData{},
	}
	for _, data := range resp.TextDataList {
		result.TextDataList = append(result.TextDataList, storage.TextDataFromGRPC(data))
	}
	for _, data := range resp.BankDataList {
		result.BankDataList = append(result.BankDataList, storage.BankDataFromGRPC(data))
	}
	for _, data := range resp.BinaryDataList {
		result.BinaryDataList = append(result.BinaryDataList, storage.BinaryDataFromGRPC(data))
	}

	return &result, nil
}
