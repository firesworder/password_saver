package mocks

import (
	"context"
	"fmt"

	"github.com/firesworder/password_saver/internal/storage"
)

type Server struct {
}

func (s *Server) RegisterUser(ctx context.Context, user storage.User) (string, error) {
	if user.Login == "admin" && user.HashedPassword == "admin" {
		return "", fmt.Errorf("test_error")
	}
	return "some_token", nil
}

func (s *Server) LoginUser(ctx context.Context, user storage.User) (string, error) {
	if user.Login == "admin" && user.HashedPassword == "admin" {
		return "", fmt.Errorf("test_error")
	}
	return "some_token", nil
}

func (s *Server) AddRecord(ctx context.Context, rawRecord interface{}) (int, error) {
	switch v := rawRecord.(type) {
	case storage.TextData:
		if v.TextData == "" {
			return 0, fmt.Errorf("test error")
		}
	case storage.BankData:
		if v.CVV == "" {
			return 0, fmt.Errorf("test error")
		}
	case storage.BinaryData:
		if v.BinaryData == nil {
			return 0, fmt.Errorf("test error")
		}
	}
	return 0, nil
}

func (s *Server) UpdateRecord(ctx context.Context, rawRecord interface{}) error {
	switch v := rawRecord.(type) {
	case storage.TextData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	case storage.BankData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	case storage.BinaryData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	}
	return nil
}

func (s *Server) DeleteRecord(ctx context.Context, rawRecord interface{}) error {
	switch v := rawRecord.(type) {
	case storage.TextData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	case storage.BankData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	case storage.BinaryData:
		if v.ID == 100 {
			return fmt.Errorf("test error")
		}
	}
	return nil
}

func (s *Server) GetAllRecords(ctx context.Context) (*storage.RecordsList, error) {
	result := &storage.RecordsList{
		TextDataList: []storage.TextData{{ID: 150, TextData: "td", MetaInfo: "mi1", UserID: 15}},
		BankDataList: []storage.BankData{
			{ID: 150, CardNumber: "1122334466778899", CardExpire: "09/12", CVV: "456", UserID: 15},
			{ID: 160, CardNumber: "0055662233661122", CardExpire: "09/24", CVV: "264", UserID: 17},
		},
		BinaryDataList: []storage.BinaryData{{ID: 200, BinaryData: []byte("demo byte"), UserID: 15}},
	}
	return result, nil
}
