package function

import (
	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddCandidat(c *gin.Context) {
	result, mess := database.AddCandidat(c.PostForm("firstname"), c.PostForm("lastname"), c.PostForm("email"), c.PostForm("formation"), c.PostForm("experience"), c.PostForm("competence"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func SearchCandidat(c *gin.Context) {
	result, mess := database.SearchCandidat(c.Query("search"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}