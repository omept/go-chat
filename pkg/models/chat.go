package models

import (
	"github.com/ong-gtp/go-chat/pkg/config"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Message    string   `json:"Message"`
	UserId     uint     `json:"UserId" gorm:"index"`
	User       User     `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChatRoomId uint     `json:"ChatRoomId" gorm:"index"`
	ChatRoom   ChatRoom `json:"ChatRoom" gorm:"foreignKey:ChatRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (cr *Chat) List(roomId uint, cht *[]Chat) *gorm.DB {
	db := config.GetDB()
	db = db.Where(Chat{ChatRoomId: roomId}).Preload("ChatRoom").Preload("User").Find(&cht).Limit(50)
	return db
}

func (c *Chat) Add() *gorm.DB {
	db := config.GetDB()
	db = db.Where(c).FirstOrCreate(&c)
	return db
}
