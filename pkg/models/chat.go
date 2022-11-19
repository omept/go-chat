package models

import "gorm.io/gorm"

type Chat struct {
	*gorm.Model
	Message    string   `json:"Message"`
	UserId     uint     `json:"UserId" gorm:"index"`
	User       User     `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChatRoomId uint     `json:"ChatRoomId" gorm:"index"`
	ChatRoom   ChatRoom `json:"ChatRoom" gorm:"foreignKey:ChatRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}
