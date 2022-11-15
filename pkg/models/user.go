package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ong-gtp/go-chat/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Uname    string  `json:"uname"`
	Password *string `json:"password,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() (*User, *gorm.DB) {
	db.NewRecord(b)
	db := db.Create(&b)
	return b, db
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", id).Find(&getUser)
	return &getUser, db
}
