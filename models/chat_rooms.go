package models

import (
	"github.com/ong-gtp/go-chat/config"
	"gorm.io/gorm"
)

type ChatRoom struct {
	gorm.Model
	Name string `json:"Name"`
}

func (cr *ChatRoom) Add() *gorm.DB {
	db := config.GetDB()
	db = db.Where(cr).FirstOrCreate(&cr)
	return db
}

func (cr *ChatRoom) List(cht *[]ChatRoom) *gorm.DB {
	db := config.GetDB()
	db = db.Order("id DESC").Find(&cht)
	return db

}
