package models

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ong-gtp/go-chat/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Uname    string `json:"uname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) GetUserByEmail(email string) *gorm.DB {
	db = db.Where("Email=?", email).Find(&u)
	return db
}

func (u *User) SaveNew() *gorm.DB {
	db = db.Create(&u)
	return db
}
