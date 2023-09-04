package storage

type User struct {
	ID                    int
	Login, HashedPassword string
}

type TextData struct {
	ID       int
	TextData string
	MetaInfo string
}

type BankData struct {
	ID         int
	CardNumber string
	CardExpire string
	CVV        string
	MetaInfo   string
}

type BinaryData struct {
	ID         int
	BinaryData []byte
	MetaInfo   string
}

type RecordsList struct {
	TextDataList   []TextData
	BankDataList   []BankData
	BinaryDataList []BinaryData
}
