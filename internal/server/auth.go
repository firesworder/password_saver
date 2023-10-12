package server

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/firesworder/password_saver/internal/storage"
)

var ErrWrongPassword = errors.New("wrong password")

const bcryptCost = 8
const genTokenSize, authTokenSize = 32, 32

// generateRandom - создает массив байт заданной длины
func generateRandom(size int) ([]byte, error) {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}
	return randBytes, nil
}

// generateToken - создает токен авторизации, с использованием hmac
func (s *Server) generateToken() ([]byte, error) {
	newToken, err := generateRandom(authTokenSize)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, s.genToken)
	h.Write(newToken)
	return h.Sum(nil), nil
}

// RegisterUser регистрирует пользователя в системе.
// Пароль формируется через bcrypt.
// Если регистрация произошла успешно - генерируется токен пользователя с записью его в хранил. токенов на сервере.
func (s *Server) RegisterUser(ctx context.Context, user storage.User) (string, error) {
	// хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcryptCost)
	if err != nil {
		return "", err
	}
	user.HashedPassword = string(hashedPassword)

	newUser, err := s.ssql.UserRep.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	uTokenBytes, err := s.generateToken()
	if err != nil {
		return "", err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers.Store(uToken, *newUser)
	return uToken, nil
}

// LoginUser авторизует пользователя в системе.
// Пароль сохраненной в БД и присланный с клиента сравниваются через bcrypt.
// Если авторизация произошла успешно - генерируется токен пользователя с записью его в хранил. токенов на сервере.
func (s *Server) LoginUser(ctx context.Context, user storage.User) (string, error) {
	bdUser, err := s.ssql.UserRep.GetUser(ctx, user)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(bdUser.HashedPassword), []byte(user.HashedPassword)); err != nil {
		return "", ErrWrongPassword
	}

	uTokenBytes, err := s.generateToken()
	if err != nil {
		return "", err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers.Store(uToken, *bdUser)
	return uToken, nil
}
