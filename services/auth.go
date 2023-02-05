package services

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ong-gtp/go-chat/http/responses"
	"github.com/ong-gtp/go-chat/models"
	"github.com/ong-gtp/go-chat/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(userName, password string) (responses.LoginResponse, error)
	SignUp(email, userName, password string) (responses.SignUpResponse, error)
}

type auth struct{}

func NewAuthService() *auth {
	return &auth{}
}

func (a *auth) Login(email, password string) (responses.LoginResponse, error) {
	var user models.User
	err := user.GetUserByEmail(email)
	if user.ID == 0 || err.Error != nil {
		return responses.LoginResponse{}, errors.ErrInvalidCredentials
	}

	jwtTTL, err1 := strconv.Atoi(os.Getenv("JWT_TTL"))
	errors.ErrorCheck(err1)
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return responses.LoginResponse{}, errors.ErrInvalidCredentials
	}

	tk := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
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

func (a *auth) SignUp(email, userName, password string) (responses.SignUpResponse, error) {
	var userCheck models.User
	userCheck.GetUserByEmail(email)
	if userCheck.ID > 0 {
		return responses.SignUpResponse{}, errors.ErrDuplicateEmail
	}

	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	errors.ErrorCheck(err)
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	errors.ErrorCheck(err)

	hPS := string(hashPass)
	user := models.User{
		Password: hPS,
		Email:    email,
		UserName: userName,
	}
	err3 := user.SaveNew()
	errors.DBErrorCheck(err3)
	user.Password = ""

	tk := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err2 := token.SignedString([]byte(jwtSecret))
	errors.ErrorCheck(err2)

	return responses.SignUpResponse{User: user, JwtToken: tokenString}, nil

}
