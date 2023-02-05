package errors

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials  = errors.New("invalid login credentials")
	ErrInRequestMarshaling = errors.New("invalid/bad request paramenters")
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrMalformedToken      = errors.New("malformed jwt token")
)

func Error(e error) {
	panic(e)
}

func DBErrorCheck(db *gorm.DB) {
	if err := db.Error; err != nil {
		Error(err)
	}
}

func ErrorCheck(e error) {
	if e != nil {
		Error(e)
	}
}
