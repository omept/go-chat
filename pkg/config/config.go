package config

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ong-gtp/go-chat/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dbUserName := os.Getenv("DB_USERNAME")
	dbUserPassword := os.Getenv("DB_PASSWORD")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dsn := dbUserName + ":" + dbUserPassword + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errors.ErrorCheck(err)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
