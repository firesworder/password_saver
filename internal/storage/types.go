package storage

// User тип объекта пользователя.
type User struct {
	ID                    int
	Login, HashedPassword string
}

// TextData тип объекта текстовых данных.
type TextData struct {
	ID       int
	TextData string
	MetaInfo string
	UserID   int
}

// BankData тип объекта банковских данных.
type BankData struct {
	ID         int
	CardNumber string
	CardExpire string
	CVV        string
	MetaInfo   string
	UserID     int
}

// BinaryData тип объекта бинарных данных.
type BinaryData struct {
	ID         int
	BinaryData []byte
	MetaInfo   string
	UserID     int
}

// RecordsList тип для метода возвр. все записи из БД(не является сущностью в БД).
// В себе хранит слайсы текстовых, банковских, и бинарных данных(пользователя).
type RecordsList struct {
	TextDataList   []TextData
	BankDataList   []BankData
	BinaryDataList []BinaryData
}

// Record формат хранения записей на сервере.
type Record struct {
	ID         int
	RecordType string
	Content    []byte
	MetaInfo   string
}
