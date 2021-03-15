package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)
// import "time"

type User struct {
	gorm.Model
	Name 				string `json:"name"`
	EmailId 			string `json:"emailid",gorm:"unique"`
	UUID 				uint `gorm:"unique"`
  }
 





