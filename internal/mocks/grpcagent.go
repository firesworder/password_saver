package mocks

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

type GrpcAgent struct {
	InputArgs []interface{}
	Users     map[string]storage.User
}

func NewGRPCAgent() *GrpcAgent {
	return &GrpcAgent{Users: map[string]storage.User{}, InputArgs: []interface{}{}}
}

func (a *GrpcAgent) RegisterUser(login, password string) error {
	a.InputArgs = []interface{}{login, password}
	if _, ok := a.Users[login+password]; ok {
		return storage.ErrLoginExist
	}
	a.Users[login+password] = storage.User{Login: login, HashedPassword: password}
	return nil
}

func (a *GrpcAgent) LoginUser(login, password string) error {
	a.InputArgs = []interface{}{login, password}
	if _, ok := a.Users[login+password]; !ok {
		return storage.ErrLoginNotExist
	}
	return nil
}

func (a *GrpcAgent) CreateTextDataRecord(input storage.TextData) (int, error) {
	a.InputArgs = []interface{}{input}
	if input.TextData == "" {
		return 0, fmt.Errorf("data invalid")
	}
	return 1, nil
}

func (a *GrpcAgent) CreateBankDataRecord(input storage.BankData) (int, error) {
	a.InputArgs = []interface{}{input}
	// input.CVV == "000" - как надежный способ стриггерить ошибку
	if input.CardNumber == "" || input.CardExpire == "" || input.CVV == "" || input.CVV == "000" {
		return 0, fmt.Errorf("data invalid")
	}
	return 1, nil
}

func (a *GrpcAgent) CreateBinaryDataRecord(input storage.BinaryData) (int, error) {
	a.InputArgs = []interface{}{input}
	if len(input.BinaryData) == 0 {
		return 0, fmt.Errorf("data invalid")
	}
	return 1, nil
}

func (a *GrpcAgent) UpdateTextDataRecord(input storage.TextData) error {
	a.InputArgs = []interface{}{input}
	if input.TextData == "" {
		return fmt.Errorf("data invalid")
	}
	return nil
}

func (a *GrpcAgent) UpdateBankDataRecord(input storage.BankData) error {
	a.InputArgs = []interface{}{input}
	// input.CVV == "000" - как надежный способ стриггерить ошибку
	if input.CardNumber == "" || input.CardExpire == "" || input.CVV == "" || input.CVV == "000" {
		return fmt.Errorf("data invalid")
	}
	return nil
}

func (a *GrpcAgent) UpdateBinaryDataRecord(input storage.BinaryData) error {
	a.InputArgs = []interface{}{input}
	if len(input.BinaryData) == 0 {
		return fmt.Errorf("data invalid")
	}
	return nil
}

func (a *GrpcAgent) DeleteTextDataRecord(input storage.TextData) error {
	a.InputArgs = []interface{}{input}
	if input.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (a *GrpcAgent) DeleteBankDataRecord(input storage.BankData) error {
	a.InputArgs = []interface{}{input}
	if input.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (a *GrpcAgent) DeleteBinaryDataRecord(input storage.BinaryData) error {
	a.InputArgs = []interface{}{input}
	if input.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (a *GrpcAgent) ShowAllRecords() (*storage.RecordsList, error) {
	return &storage.RecordsList{
		TextDataList: []storage.TextData{
			{ID: 1, TextData: "Aranara", MetaInfo: "meta info", UserID: 1},
			{ID: 10, TextData: "Ararakalari", MetaInfo: "meta info", UserID: 1},
		},
		BankDataList: []storage.BankData{
			{ID: 2, CardNumber: "0011223344556677", CardExpire: "12/23", CVV: "453", MetaInfo: "meta info", UserID: 1},
		},
		BinaryDataList: []storage.BinaryData{
			{ID: 3, BinaryData: []byte("Aranara"), MetaInfo: "meta info", UserID: 1},
		},
	}, nil
}
