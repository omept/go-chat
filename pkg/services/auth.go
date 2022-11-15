package services

import (
	"github.com/ong-gtp/go-chat/pkg/domain/responses"
)

type AuthService interface {
	Login(uname, password string) (responses.LoginResponse, error)
	SignUp(uname, password string) (responses.SignUpResponse, error)
}

type auth struct{}

func NewAuthService() AuthService {
	return &auth{}
}

func (a *auth) Login(uname, password string) (responses.LoginResponse, error) {
	return responses.LoginResponse{}, nil
}

func (a *auth) SignUp(uname, password string) (responses.SignUpResponse, error) {
	return responses.SignUpResponse{}, nil
}
