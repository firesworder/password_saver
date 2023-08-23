package storage

type User struct {
	Login, HashedPassword string
}

type MetaInfo struct {
	ID   int
	Info string
}

type TextData struct {
	ID       int
	TextData string
}
