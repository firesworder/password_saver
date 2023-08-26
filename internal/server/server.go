package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/mocks/users"
)

var ErrWrongPassword = errors.New("wrong password")

type Server struct {
	env       *env.Environment
	uRep      storage.UserRepository
	authUsers map[string]storage.User
}

func NewServer() (*Server, error) {
	var uRep storage.UserRepository
	// todo: удалить после разработки
	if true {
		uRep = &users.MockUser{Users: map[string]storage.User{}}
	}
	s := &Server{env: &env.Env, uRep: uRep, authUsers: map[string]storage.User{}}
	return s, nil
}

// generateRandom - создает массив байт заданной длины
func generateRandom(size int) ([]byte, error) {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}
	return randBytes, nil
}

func generateToken() ([]byte, error) {
	return generateRandom(32)
}

func (s *Server) RegisterUser(ctx context.Context, user storage.User) error {
	// хеширование пароля

	err := s.uRep.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	uTokenBytes, err := generateToken()
	if err != nil {
		return err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers[uToken] = user
	return nil
}

func (s *Server) LoginUser(ctx context.Context, user storage.User) error {
	bdUser, err := s.uRep.GetUser(ctx, user)
	if err != nil {
		return err
	}
	if bdUser.HashedPassword != user.HashedPassword {
		return ErrWrongPassword
	}

	uTokenBytes, err := generateToken()
	if err != nil {
		return err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers[uToken] = user
	return nil
}
