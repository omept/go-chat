package services

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ong-gtp/go-chat/pkg/domain/responses"
	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/models"
	"github.com/ong-gtp/go-chat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(uname, password string) (responses.LoginResponse, error)
	SignUp(uname, password string) (responses.SignUpResponse, error)
}

type auth struct{}

func NewAuthService() *auth {
	return &auth{}
}

func (a *auth) Login(email, password string) (responses.LoginResponse, error) {
	user, err := models.GetUserByEmail(email)
	if user.ID == 0 || err != nil {
		return responses.LoginResponse{}, errors.ErrInvalidCredentials
	}

	jwtTTL, err1 := strconv.Atoi(os.Getenv("JWT_TTL"))
	errors.ErrorCheck(err1)
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(utils.DerefString(user.Password)), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return responses.LoginResponse{}, errors.ErrInvalidCredentials
	}

	tk := &models.Token{
		UserID: user.ID,
		Uname:  user.Uname,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err2 := token.SignedString([]byte(jwtSecret))
	errors.ErrorCheck(err2)

	return responses.LoginResponse{User: user, JwtToken: tokenString}, nil
}

func (a *auth) SignUp(uname, password string) (responses.SignUpResponse, error) {
	return responses.SignUpResponse{}, nil
}
