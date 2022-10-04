package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	result, mess := database.Register(c.PostForm("username"), c.PostForm("email"), c.PostForm("password"))
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

func UpdateRole(c *gin.Context) {
	username := c.PostForm("username")
	role, err := strconv.Atoi(c.PostForm("role"))
	if err != nil {
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
