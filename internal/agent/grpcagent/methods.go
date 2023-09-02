package grpcagent

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
)

func (a *GRPCAgent) RegisterUser(login, password string) (string, error) {
	req := pb.RegisterUserRequest{Login: login, Password: password}

	resp, err := a.grpcClient.RegisterUser(context.Background(), &req)
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (a *GRPCAgent) LoginUser(login, password string) (string, error) {
	req := pb.LoginUserRequest{Login: login, Password: password}

	resp, err := a.grpcClient.LoginUser(context.Background(), &req)
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (a *GRPCAgent) CreateTextDataRecord(input storage.TextData) (int, error) {
	req := pb.AddTextDataRequest{TextData: storage.TextDataToGRPC(input)}

	resp, err := a.grpcClient.AddTextDataRecord(context.Background(), &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

func (a *GRPCAgent) CreateBankDataRecord(input storage.BankData) (int, error) {
	req := pb.AddBankDataRequest{BankData: storage.BankDataToGRPC(input)}

	resp, err := a.grpcClient.AddBankDataRecord(context.Background(), &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

func (a *GRPCAgent) CreateBinaryDataRecord(input storage.BinaryData) (int, error) {
	req := pb.AddBinaryDataRequest{BinaryData: storage.BinaryDataToGRPC(input)}

	resp, err := a.grpcClient.AddBinaryDataRecord(context.Background(), &req)
	if err != nil {
		return 0, err
	}

	return int(resp.Id), nil
}

func (a *GRPCAgent) UpdateTextDataRecord(input storage.TextData) error {
	req := pb.UpdateTextDataRequest{TextData: storage.TextDataToGRPC(input)}

	_, err := a.grpcClient.UpdateTextDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) UpdateBankDataRecord(input storage.BankData) error {
	req := pb.UpdateBankDataRequest{BankData: storage.BankDataToGRPC(input)}

	_, err := a.grpcClient.UpdateBankDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) UpdateBinaryDataRecord(input storage.BinaryData) error {
	req := pb.UpdateBinaryDataRequest{BinaryData: storage.BinaryDataToGRPC(input)}

	_, err := a.grpcClient.UpdateBinaryDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) DeleteTextDataRecord(input storage.TextData) error {
	req := pb.DeleteTextDataRequest{Id: int64(input.ID)}

	_, err := a.grpcClient.DeleteTextDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) DeleteBankDataRecord(input storage.BankData) error {
	req := pb.DeleteBankDataRequest{Id: int64(input.ID)}

	_, err := a.grpcClient.DeleteBankDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) DeleteBinaryDataRecord(input storage.BinaryData) error {
	req := pb.DeleteBinaryDataRequest{Id: int64(input.ID)}

	_, err := a.grpcClient.DeleteBinaryDataRecord(context.Background(), &req)
	return err
}

func (a *GRPCAgent) ShowAllRecords() (*storage.RecordsList, error) {
	req := pb.GetAllRecordsRequest{}

	resp, err := a.grpcClient.GetAllRecords(context.Background(), &req)
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
