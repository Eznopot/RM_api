package function

import (
	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	result, mess := database.Login(c.PostForm("username"), c.PostForm("password"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func ChangePassword(c *gin.Context) {
	result, mess := database.ChangePassword(c.Request.Header["Token"][0], c.PostForm("password"), c.PostForm("newpassword"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func Logout(c *gin.Context) {
	result, mess := database.Logout(c.Request.Header["Token"][0])
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func GetInfo(c *gin.Context) {
	result, mess := database.GetInfo(c.Request.Header["Token"][0])
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func GetPages(c *gin.Context) {
	result, mess := database.GetPages(c.Request.Header["Token"][0])
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func CheckSession(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Session is valid",
		"result":  true,
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test effectued",
		"result":  "Hello World",
	})
}
