package server

import (
	"github.com/firesworder/password_saver/internal/server/env"
)

type Server struct {
	env *env.Environment
}

func NewServer() (*Server, error) {
	s := &Server{env: &env.Env}
	return s, nil
}
