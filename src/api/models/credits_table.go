package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserCredit struct {
	ID				uint `gorm:"primarykey"`
	Credited 		uint 
	User User `gorm:"foreignkey:UserId"`
	UserId    		uint `gorm:"size:10"`
	UUID uint
	Priority uint
	Subscription string
	Expiry string
	Debited uint
  }