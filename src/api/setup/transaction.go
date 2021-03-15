package Setup

import(
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"api/models"
	"api/connection"
	"api/structs"
)

func CreditsAccount(c *gin.Context) {
	var input structs.Credits
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusOK, gin.H{"Incorrect JSON": err})
	return
}

if input.Activity == "credit"{
	var user models.User
	var useraccountbal models.UserAccountBalance
	connection.DB.First(&user, "uuid = ?", input.Key.Uuid)
	if user.ID != 0 {
		credit := models.UserCredit{Credited: input.Key.Amount, Priority: input.Key.Priority, Subscription: input.Key.Type, Expiry:input.Key.Expiry,UUID:input.Key.Uuid,UserId:user.ID}
		connection.DB.Create(&credit)
		connection.DB.First(&useraccountbal, "user_id = ?",user.ID)
		creditedAmt := useraccountbal.Balance + input.Key.Amount
		connection.DB.Table("user_account_balances").Where("user_id = ?",user.ID ).Update("balance",creditedAmt)
		connection.DB.First(&useraccountbal, "user_id = ?",user.ID)
		c.JSON(http.StatusOK, gin.H{"Amount credited": input.Key.Amount,"Your Balance is":useraccountbal.Balance})
	}else{
		c.JSON(http.StatusOK, gin.H{"User does not exist":"PLease try with proper UUID"})
	}
}

}

