package function

import (
	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	result, mess := database.Register(c.PostForm("username"), c.PostForm("password"));
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}

func Login(c *gin.Context) {
	result, mess := database.Login(c.PostForm("username"), c.PostForm("password"));
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}

func Logout(c *gin.Context) {
	result, mess := database.Logout(c.Request.Header["Token"][0]);
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}

func GetInfo(c *gin.Context) {
	result, mess := database.GetInfo(c.Request.Header["Token"][0]);
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}

func GetAllUser(c *gin.Context) {
	result, mess := database.GetAllUser();
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}

func GetRoutes(c *gin.Context) {
	result, mess := database.GetRoutes();
	c.JSON(200, gin.H{
		"message": mess,
		"result": result,
	});
}