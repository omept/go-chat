package models

import (
	"github.com/ong-gtp/go-chat/pkg/config"
	"gorm.io/gorm"
)

// var db *gorm.DB

type User struct {
	gorm.Model
	UserName string `json:"UserName,omitempty"`
	Email    string `json:"Email,omitempty"`
	Password string `json:"Password,omitempty"`
}

func (u *User) GetUserByEmail(email string) *gorm.DB {
	db := config.GetDB()
	db = db.Where("Email=?", email).Find(&u)
	return db
}

func (u *User) SaveNew() *gorm.DB {
	db := config.GetDB()
	db = db.Create(&u)
	return db
}
