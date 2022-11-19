package models

import (
	"github.com/ong-gtp/go-chat/pkg/config"
	"gorm.io/gorm"
)

type ChatRoom struct {
	*gorm.Model
	Name string `json:"Name"`
}

func (cr *ChatRoom) Add() *gorm.DB {
	db := config.GetDB()
	db = db.FirstOrCreate(&cr)
	return db
}
