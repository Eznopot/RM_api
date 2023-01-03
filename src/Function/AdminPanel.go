package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{
			"message": "Bad argument",
			"result":  false,
		})
	}
	result, mess := database.Register(id, c.PostForm("username"), c.PostForm("firstname"), c.PostForm("lastname"), c.PostForm("phone"), c.PostForm("email"), c.PostForm("password"), c.PostForm("address"), c.PostForm("postalCode"), c.PostForm("country"), c.PostForm("emergencyName"), c.PostForm("emergencyPhonePerso"), c.PostForm("emergencyPhonePro"), c.PostForm("emergencyLink"), c.PostForm("emergencyAddress"), c.PostForm("emergencyPostalCode"), c.PostForm("emergencyCountry"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func GetAllUser(c *gin.Context) {
	result, mess := database.GetAllUser()
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func GetUserCv(c *gin.Context) {
	result, mess := database.GetUserCv(c.Query("email"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func UpdateRole(c *gin.Context) {
	username := c.PostForm("username")
	role, err := strconv.Atoi(c.PostForm("role"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{
			"message": "Bad argument",
			"result":  false,
		})
	}
	result, mess := database.UpdateRole(username, role)
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func AddAdminString(c *gin.Context) {
	result, mess := database.AddAdminString(c.PostForm("value"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func ModifyAdminString(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{
			"message": "Bad argument",
			"result":  false,
		})
	}
	result, mess := database.ModifyAdminString(id, c.PostForm("value"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func DeleteAdminString(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{
			"message": "Bad argument",
			"result":  false,
		})
	}
	result, mess := database.DeleteAdminString(id)
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func GetAdminString(c *gin.Context) {
	result, mess := database.GetAdminString()
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}
