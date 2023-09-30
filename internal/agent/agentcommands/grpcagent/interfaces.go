package grpcagent

import "github.com/firesworder/password_saver/internal/storage"

// IGRPCAgent Интерфейс grpc-агента. (основное назначение - подмена моком, при тестировании)
type IGRPCAgent interface {
	RegisterUser(login, password string) error
	LoginUser(login, password string) error
	CreateTextDataRecord(input storage.TextData) (int, error)
	CreateBankDataRecord(input storage.BankData) (int, error)
	CreateBinaryDataRecord(input storage.BinaryData) (int, error)
	UpdateTextDataRecord(input storage.TextData) error
	UpdateBankDataRecord(input storage.BankData) error
	UpdateBinaryDataRecord(input storage.BinaryData) error
	DeleteTextDataRecord(input storage.TextData) error
	DeleteBankDataRecord(input storage.BankData) error
	DeleteBinaryDataRecord(input storage.BinaryData) error
	ShowAllRecords() (*storage.RecordsList, error)
}
