package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserDebit struct {
	ID				uint `gorm:"primarykey"`
	Debited 		uint 
	User User `gorm:"foreignkey:UserId"`
	UserId    		uint `gorm:"size:10"`
  }