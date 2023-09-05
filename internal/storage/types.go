package storage

type User struct {
	ID                    int
	Login, HashedPassword string
}

type TextData struct {
	ID       int
	TextData string
	MetaInfo string
	UserID   int
}

type BankData struct {
	ID         int
	CardNumber string
	CardExpire string
	CVV        string
	MetaInfo   string
	UserID     int
}

type BinaryData struct {
	ID         int
	BinaryData []byte
	MetaInfo   string
	UserID     int
}

type RecordsList struct {
	TextDataList   []TextData
	BankDataList   []BankData
	BinaryDataList []BinaryData
}
