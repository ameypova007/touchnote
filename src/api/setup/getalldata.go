package Setup

import(
	// "fmt"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"api/models"
	"api/connection"
	// "api/structs"
)

func GetAllData(c *gin.Context) {
	var credits []models.UserCredit
	var debits []models.UserDebit

	connection.DB.Find(&credits)
	connection.DB.Find(&debits)

	c.JSON(http.StatusOK, gin.H{"credits": credits,"debits":debits})
}


func SpecificUserData(c *gin.Context) {
	var credits []models.UserCredit
	var debits []models.UserDebit
	var user models.User

	connection.DB.First(&user, "uuid = ?", c.Param("uuid"))

	connection.DB.Where("user_id = ?", user.ID).Find(&credits)
	if err := connection.DB.Where("user_id = ?", user.ID).Find(&credits).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := connection.DB.Where("user_id = ?", user.ID).Find(&debits).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"credits": credits,"debits":debits})
}