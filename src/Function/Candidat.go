package function

import (
	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/Eznopot/RM_api/src/utils"
	"github.com/gin-gonic/gin"
)

func AddCandidat(c *gin.Context) {
	result, mess, id := database.AddCandidat(c.PostForm("firstname"), c.PostForm("lastname"), c.PostForm("email"), c.PostForm("phone"), c.PostForm("formations"), c.PostForm("experiences"), c.PostForm("competences"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
	utils.CreateWordCV(id, c.PostForm("firstname")[0:1]+"."+c.PostForm("lastname")[0:1], c.PostForm("formation"), c.PostForm("experience"), c.PostForm("competence"))
}

func SearchCandidat(c *gin.Context) {
	result, mess := database.SearchCandidat(c.Query("search"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func SearchCandidatByEmail(c *gin.Context) {
	result, mess := database.SearchCandidatByEmail(c.Query("email"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}
