package main

import (
	// "fmt"
	// "log"
	// "github.com/gorilla/mux" 
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	// "api/models"
	"api/connection"
	"api/setup"
	// "github.com/jasonlvhit/gocron"
	// "api/setup"
	// "api/Structs"
	// "time"
)


func main() {
   
	r := gin.Default()
	// db.Migrator()
	connection.ConnectDatabase()
	

	r.POST("/hello", Setup.CreateAccount)
	r.POST("/credits", Setup.CreditsAccount)
	r.POST("/debits",Setup.DebitsAccount)
	r.GET("/getAll",Setup.GetAllData)
	r.GET("get/:uuid",Setup.SpecificUserData)
	// gocron.Every(1).Second().Do(Setup.Task)
	r.Run()
}