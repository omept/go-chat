package errors

import (
	"errors"

	"github.com/jinzhu/gorm"
	// "github.com/ong-gtp/go-chat/pkg/domain/responses"
)

var (
	ErrInvalidCredentials = errors.New("invalid login credentials")
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
