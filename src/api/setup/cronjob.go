package Setup

// import (
// 	"fmt"
// 	"time"
// 	"api/models"
// 	"api/connection"
// 	// "api/structs"
// 	"strconv"

// )
// func Task() {
// 	fmt.Println("cron started")
// 	var useraccountbal models.UserAccountBalance
// 	var user models.User
// 	var credits []models.UserCredit
// 	var credits1 models.UserCredit	
// 	if credits1.ID != 0{
// 	for i, s := range credits {
// 			if s.Debited != 1{
// 				//check what is the date of expiry
// 				convertedtime, err := strconv.Atoi(s.Expiry)
// 				fmt.Println(i,err)
// 				now := time.Now().Unix() 
// 				expiryTime := convertedtime + int(now)
// 				if expiryTime <= int(now) {
// 				connection.DB.First(&useraccountbal, "user_id = ?",s.ID)
// 				fmt.Println(useraccountbal.Balance)
// 				amountremain := useraccountbal.Balance - s.Credited
// 				connection.DB.Table("user_account_balances").Where("user_id = ?",s.ID).Update("balance",amountremain)
// 				connection.DB.First(&credits1, "user_id = ?",user.ID)
// 				fmt.Println(credits1)
// 				connection.DB.Table("user_credits").Where("user_id = ? ",s.ID).Update("debited",1)
// 				debit := models.UserDebit{Debited:s.Credited,UserId:user.ID}
// 				connection.DB.Create(&debit)
// 				}
// 		}
	
// }
// }
// }