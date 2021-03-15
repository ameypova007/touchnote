package connection

import (
	// "github.com/gorilla/mux" 
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"api/models"
	// "time"
	"log"
)
var DB *gorm.DB
var err error

func ConnectDatabase() {
database, err := gorm.Open("mysql", "root:i13D04ms$12@tcp(127.0.0.1:3306)/sample?charset=utf8&parseTime=True")
	if err!=nil{
	log.Println("Connection Failed to Open",err)
	}else{ 
	log.Println("Connection Established")
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.UserAccountBalance{})
	database.AutoMigrate(&models.UserCredit{})
	database.AutoMigrate(&models.UserDebit{})
	database.Table("user_account_balances").AddForeignKey("user_id","users(id)","CASCADE", "CASCADE")
	database.Table("user_debits").AddForeignKey("user_id","users(id)","CASCADE", "CASCADE")
	database.Table("user_credits").AddForeignKey("user_id","users(id)","CASCADE", "CASCADE")
	DB = database
}