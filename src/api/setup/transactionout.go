package Setup

import(
	"fmt"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"api/models"
	"api/connection"
	"api/structs"
)

func DebitsAccount(c *gin.Context) {
	var input structs.Debits
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusOK, gin.H{"Incorrect JSON": err})
	return
}
if input.Activity == "debit"{
	var user models.User
	var useraccountbal models.UserAccountBalance
	var credits []models.UserCredit
	connection.DB.First(&user, "uuid = ?", input.Key.Uuid)
	if  user.ID != 0{
		connection.DB.First(&useraccountbal, "user_id = ?",user.ID)
		if useraccountbal.Balance < input.Key.Amount{
			c.JSON(http.StatusOK, gin.H{"Insufficient Balance, Your current balance is ": useraccountbal.Balance})
			return
		}
		connection.DB.Order("priority desc").Where("user_id = ?",user.ID).Find(&credits)
		var amt1 uint
		amt := input.Key.Amount
		for i, s := range credits {
			fmt.Print(i)
			if amt != 0{
				if s.Debited != 1 && s.Credited >= amt{
				amt1 = s.Credited - amt
				connection.DB.Table("user_credits").Where("user_id = ? AND priority = ?",user.ID,s.Priority).Update("credited",amt1)
				amt = 0
				}else if s.Debited != 1 && s.Credited < amt{
				connection.DB.Table("user_credits").Where("user_id = ? AND priority = ?",user.ID,s.Priority).Update("debited",1)
				amt = input.Key.Amount - s.Credited
				connection.DB.Table("user_credits").Where("user_id = ? AND priority = ?",user.ID,s.Priority).Update("credited",0)
				
			}
		}
		}
		connection.DB.Where("user_id = ?",user.ID).Find(&credits)
		for i, s := range credits {
			fmt.Print(i)
			if s.Credited == 0 && s.Debited!=1 {
				connection.DB.Table("user_credits").Where("user_id = ?",s.UserId).Update("debited",1)
			}
		}

		debitedAmt := useraccountbal.Balance - input.Key.Amount
		connection.DB.Table("user_account_balances").Where("user_id = ?",user.ID).Update("balance",debitedAmt)
		connection.DB.First(&useraccountbal, "user_id = ?",user.ID)
		c.JSON(http.StatusOK, gin.H{"Amount debited": input.Key.Amount,"Your Balance is":useraccountbal.Balance})
		debit := models.UserDebit{Debited: input.Key.Amount,UserId:user.ID}
		connection.DB.Create(&debit)
	}else{
		c.JSON(http.StatusOK, gin.H{"User Doesn't Exist": "PLease enter correct UUID"})
	}
}
}