package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ong-gtp/go-chat/pkg/errors"
)

var (
	db *gorm.DB
)

func ConnectDB() {

	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	d, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True")

	errors.ErrorCheck(err)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
