package models

import (
	"github.com/ong-gtp/go-chat/pkg/config"
	"gorm.io/gorm"
)

type ChartRoom struct {
	*gorm.Model
	Name string `json:"Name"`
}

func (cr *ChartRoom) Add() *gorm.DB {
	db := config.GetDB()
	db = db.Create(&cr)
	return db
}
