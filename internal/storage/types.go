package storage

type User struct {
	Login, HashedPassword string
}

type TextData struct {
	ID       int
	TextData string
	MetaInfo string
}
