package Setup

import(
	// "fmt"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"api/models"
	"api/connection"
	"api/structs"
)

func CreateAccount(c *gin.Context) {
	var input structs.Createuser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"Incorrect JSON": err})
	return
}

user := models.User{Name: input.Name, EmailId: input.EmailId, UUID:input.Uuid}
if err := connection.DB.Where("email_id = ?",input.EmailId).First(&models.User{}).Error; err != nil {
	connection.DB.Create(&user)
	useraccount := models.UserAccountBalance{Balance:0,UserId:user.ID}
	connection.DB.Create(&useraccount)
	c.JSON(http.StatusOK, gin.H{"User has been created": user})
}else{
	c.JSON(http.StatusOK, gin.H{"User is already Exists": user})
}
}