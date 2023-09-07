package server

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var ErrWrongPassword = errors.New("wrong password")

const bcryptCost = 8

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
	newToken, err := generateRandom(32)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, s.genToken)
	h.Write(newToken)
	return h.Sum(nil), err
}

func (s *Server) RegisterUser(ctx context.Context, user storage.User) (string, error) {
	// хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcryptCost)
	if err != nil {
		return "", err
	}
	user.HashedPassword = string(hashedPassword)

	newUser, err := s.uRep.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	uTokenBytes, err := s.generateToken()
	if err != nil {
		return "", err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers[uToken] = *newUser
	return uToken, nil
}

func (s *Server) LoginUser(ctx context.Context, user storage.User) (string, error) {
	// хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcryptCost)
	if err != nil {
		return "", err
	}
	user.HashedPassword = string(hashedPassword)

	bdUser, err := s.uRep.GetUser(ctx, user)
	if err != nil {
		return "", err
	}
	if bdUser.HashedPassword != user.HashedPassword {
		return "", ErrWrongPassword
	}

	uTokenBytes, err := s.generateToken()
	if err != nil {
		return "", err
	}
	uToken := hex.EncodeToString(uTokenBytes)

	s.authUsers[uToken] = user
	return uToken, nil
}
