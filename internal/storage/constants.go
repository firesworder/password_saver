package storage

import "errors"

const DevDSN = "postgresql://postgres:admin@localhost:5432/password_saver"

var (
	ErrLoginExist    = errors.New("login already exist")
	ErrLoginNotExist = errors.New("login not exist")
)
