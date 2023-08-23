package storage

type User struct {
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
