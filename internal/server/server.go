package server

import (
	"context"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/mocks/bankdata"
	"github.com/firesworder/password_saver/internal/storage/mocks/binarydata"
	"github.com/firesworder/password_saver/internal/storage/mocks/textdata"
	"github.com/firesworder/password_saver/internal/storage/mocks/users"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
)

type Server struct {
	env       *env.Environment
	uRep      storage.UserRepository
	authUsers map[string]storage.User
	tRep      storage.TextDataRepository
	bankRep   storage.BankDataRepository
	binRep    storage.BinaryDataRepository

	genToken []byte
}

func NewServer() (*Server, error) {
	var uRep storage.UserRepository
	var tRep storage.TextDataRepository
	var bankRep storage.BankDataRepository
	var binRep storage.BinaryDataRepository
	// todo: удалить после разработки
	if true {
		uRep = &users.MockUser{Users: map[string]storage.User{}}
		tRep = &textdata.MockTextData{TextDataMap: map[int]storage.TextData{}}
		bankRep = &bankdata.MockBankData{BankData: map[int]storage.BankData{}}
		binRep = &binarydata.MockBinaryData{BinaryData: map[int]storage.BinaryData{}}
	} else {
		ssql, err := sqlstorage.NewStorage(storage.DevDSN)
		if err != nil {
			return nil, err
		}
		uRep = ssql.UserRep
		tRep = ssql.TextRep
		bankRep = ssql.BankRep
		binRep = ssql.BinaryRep
	}

	genToken, err := generateRandom(32)
	if err != nil {
		return nil, err
	}
	s := &Server{
		env:       &env.Env,
		uRep:      uRep,
		tRep:      tRep,
		bankRep:   bankRep,
		binRep:    binRep,
		authUsers: map[string]storage.User{},
		genToken:  genToken,
	}
	return s, nil
}

func (s *Server) AddTextData(ctx context.Context, textData storage.TextData) (int, error) {
	id, err := s.tRep.AddTextData(ctx, textData)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateTextData(ctx context.Context, textData storage.TextData) error {
	err := s.tRep.UpdateTextData(ctx, textData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteTextData(ctx context.Context, textData storage.TextData) error {
	err := s.tRep.DeleteTextData(ctx, textData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) AddBankData(ctx context.Context, bankData storage.BankData) (int, error) {
	id, err := s.bankRep.AddBankData(ctx, bankData)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateBankData(ctx context.Context, bankData storage.BankData) error {
	err := s.bankRep.UpdateBankData(ctx, bankData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteBankData(ctx context.Context, bankData storage.BankData) error {
	err := s.bankRep.DeleteBankData(ctx, bankData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) AddBinaryData(ctx context.Context, binaryData storage.BinaryData) (int, error) {
	id, err := s.binRep.AddBinaryData(ctx, binaryData)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	err := s.binRep.UpdateBinaryData(ctx, binaryData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	err := s.binRep.DeleteBinaryData(ctx, binaryData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetAllRecords(ctx context.Context) (*storage.RecordsList, error) {
	var err error
	recList := &storage.RecordsList{}
	recList.TextDataList, err = s.tRep.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}
	recList.BankDataList, err = s.bankRep.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}
	recList.BinaryDataList, err = s.binRep.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}
	return recList, nil
}
