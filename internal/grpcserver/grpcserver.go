// Package grpcserver реализует grpc сервер по "контракту" PasswordSaverServer описанному в .proto файле.
// Основной тип пакета GRPCService - реализует функционал пакета, представляет собой grpc обертку над серверной
// логикой и соотв-но имплементацией pb.UnimplementedPasswordSaverServer.
package grpcserver

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
)

// GRPCService экземпляр grpc сервиса для запуска в grpc.NewServer.
// В себе хранит интерфейс pb.UnimplementedPasswordSaverServer и переменную serv - ссылку на объект пакета server
// с серверной функц.
type GRPCService struct {
	pb.UnimplementedPasswordSaverServer

	serv server.IServer
}

// NewGRPCService конструктор grpc сервера(обертка над server.Server).
func NewGRPCService(s server.IServer) (*GRPCService, error) {
	grpcService := &GRPCService{serv: s}
	return grpcService, nil
}

// PrepareServer создает экземпляр grpc сервера, на основе реализ. в этом пакете grpc сервиса и сертификатов TLS
// для обеспеч. защищенного соединения.
func (gs *GRPCService) PrepareServer(env *env.Environment) (*grpc.Server, error) {
	creds, err := credentials.NewServerTLSFromFile(env.CertFile, env.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	serverGRPC := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterPasswordSaverServer(serverGRPC, gs)
	return serverGRPC, nil
}

// RegisterUser регистрирует пользователя.
// Если переданы пустой логин или пароль - возвращает ошибку.
func (gs *GRPCService) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password fields can not be empty")
	}

	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}
	userToken, err := gs.serv.RegisterUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.RegisterUserResponse{Token: userToken}
	return &resp, nil
}

// LoginUser авторизует пользователя.
// Если переданы пустой логин или пароль - возвращает ошибку.
func (gs *GRPCService) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password fields can not be empty")
	}

	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}
	userToken, err := gs.serv.LoginUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.LoginUserResponse{Token: userToken}
	return &resp, nil
}

// AddTextDataRecord создает текстовую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest) (*pb.AddTextDataResponse, error) {
	id, err := gs.serv.AddRecord(ctx, storage.TextData{
		TextData: request.TextData.TextData,
		MetaInfo: request.TextData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddTextDataResponse{Id: int64(id)}
	return resp, nil
}

// UpdateTextDataRecord обновляет текстовую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	err := gs.serv.UpdateRecord(ctx, storage.TextData{
		ID:       int(request.TextData.Id),
		TextData: request.TextData.TextData,
		MetaInfo: request.TextData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateTextDataResponse{}
	return resp, nil
}

// DeleteTextDataRecord удаляет текстовую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	err := gs.serv.DeleteRecord(ctx, storage.TextData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteTextDataResponse{}
	return resp, nil
}

// AddBankDataRecord добавляет банковскую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
	id, err := gs.serv.AddRecord(ctx, storage.BankData{
		CardNumber: request.BankData.CardNumber,
		CardExpire: request.BankData.CardExpiry,
		CVV:        request.BankData.Cvv,
		MetaInfo:   request.BankData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddBankDataResponse{Id: int64(id)}
	return resp, nil
}

// UpdateBankDataRecord обновляет банковскую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
	err := gs.serv.UpdateRecord(ctx, storage.BankData{
		ID:         int(request.BankData.Id),
		CardNumber: request.BankData.CardNumber,
		CardExpire: request.BankData.CardExpiry,
		CVV:        request.BankData.Cvv,
		MetaInfo:   request.BankData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateBankDataResponse{}
	return resp, nil
}

// DeleteBankDataRecord удаляет банковскую запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	err := gs.serv.DeleteRecord(ctx, storage.BankData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBankDataResponse{}
	return resp, nil
}

// AddBinaryDataRecord добавляет бинарную запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
	id, err := gs.serv.AddRecord(ctx, storage.BinaryData{
		BinaryData: request.BinaryData.BinaryData,
		MetaInfo:   request.BinaryData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddBinaryDataResponse{Id: int64(id)}
	return resp, nil
}

// UpdateBinaryDataRecord обновляет бинарную запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
	err := gs.serv.UpdateRecord(ctx, storage.BinaryData{
		ID:         int(request.BinaryData.Id),
		BinaryData: request.BinaryData.BinaryData,
		MetaInfo:   request.BinaryData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateBinaryDataResponse{}
	return resp, nil
}

// DeleteBinaryDataRecord удаляет бинарную запись данных.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	err := gs.serv.DeleteRecord(ctx, storage.BinaryData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBinaryDataResponse{}
	return resp, nil
}

// GetAllRecords возвращает все записи пользователя.
// Требуется в контексте передать токен пользователя!
func (gs *GRPCService) GetAllRecords(ctx context.Context, request *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {
	records, err := gs.serv.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}

	textDL := make([]*pb.TextData, 0, len(records.TextDataList))
	bankDL := make([]*pb.BankData, 0, len(records.BankDataList))
	binaryDL := make([]*pb.BinaryData, 0, len(records.BinaryDataList))

	for _, v := range records.TextDataList {
		textDL = append(textDL, &pb.TextData{
			Id:       int64(v.ID),
			TextData: v.TextData,
			MetaInfo: v.MetaInfo,
		})
	}

	for _, v := range records.BankDataList {
		bankDL = append(bankDL, &pb.BankData{
			Id:         int64(v.ID),
			CardNumber: v.CardNumber,
			CardExpiry: v.CardExpire,
			Cvv:        v.CVV,
			MetaInfo:   v.MetaInfo,
		})
	}

	for _, v := range records.BinaryDataList {
		binaryDL = append(binaryDL, &pb.BinaryData{
			Id:         int64(v.ID),
			BinaryData: v.BinaryData,
			MetaInfo:   v.MetaInfo,
		})
	}

	resp := &pb.GetAllRecordsResponse{
		TextDataList:   textDL,
		BankDataList:   bankDL,
		BinaryDataList: binaryDL,
	}
	return resp, nil
}
