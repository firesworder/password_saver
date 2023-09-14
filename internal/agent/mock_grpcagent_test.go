package agent

import (
	"github.com/firesworder/password_saver/internal/storage"
)

type MockGRPCAgent struct {
	inputArgs []interface{}
}

func (m *MockGRPCAgent) RegisterUser(login, password string) error {
	m.inputArgs = []interface{}{login, password}
	return nil
}

func (m *MockGRPCAgent) LoginUser(login, password string) error {
	m.inputArgs = []interface{}{login, password}
	return nil
}

func (m *MockGRPCAgent) CreateTextDataRecord(input storage.TextData) (int, error) {
	m.inputArgs = []interface{}{input}
	return 1, nil
}

func (m *MockGRPCAgent) CreateBankDataRecord(input storage.BankData) (int, error) {
	m.inputArgs = []interface{}{input}
	return 1, nil
}

func (m *MockGRPCAgent) CreateBinaryDataRecord(input storage.BinaryData) (int, error) {
	m.inputArgs = []interface{}{input}
	return 1, nil
}

func (m *MockGRPCAgent) UpdateTextDataRecord(input storage.TextData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) UpdateBankDataRecord(input storage.BankData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) UpdateBinaryDataRecord(input storage.BinaryData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) DeleteTextDataRecord(input storage.TextData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) DeleteBankDataRecord(input storage.BankData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) DeleteBinaryDataRecord(input storage.BinaryData) error {
	m.inputArgs = []interface{}{input}
	return nil
}

func (m *MockGRPCAgent) ShowAllRecords() (*storage.RecordsList, error) {
	return &storage.RecordsList{
		TextDataList: []storage.TextData{
			{ID: 1, TextData: "Aranara", MetaInfo: "meta info", UserID: 1},
		},
		BankDataList: []storage.BankData{
			{ID: 1, CardNumber: "0011223344556677", CardExpire: "12/23", CVV: "453", MetaInfo: "meta info", UserID: 1},
		},
		BinaryDataList: []storage.BinaryData{
			{ID: 1, BinaryData: []byte("Aranara"), MetaInfo: "meta info", UserID: 1},
		},
	}, nil
}
