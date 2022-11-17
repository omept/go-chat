package models

import "gorm.io/gorm"

type Chart struct {
	*gorm.Model
	UserId      uint      `json:"UserId" gorm:"index"`
	User        User      `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChartRoomId uint      `json:"ChartRoomId" gorm:"index"`
	ChartRoom   ChartRoom `json:"ChartRoom" gorm:"foreignKey:ChartRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}
