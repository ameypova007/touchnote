package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserAccountBalance struct {
	ID				uint `gorm:"primarykey"`
	Balance 		uint `json:"balance"`
	User User `gorm:"foreignkey:UserId"`
	UserId    		uint `gorm:"size:10"`
  }