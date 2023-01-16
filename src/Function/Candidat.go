package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddCandidat(c *gin.Context) {
	offerId, err := strconv.Atoi(c.PostForm("offerId"))
	if (err != nil) {
		offerId = 0
	}
	result, mess := database.AddCandidat(c.PostForm("firstname"), c.PostForm("lastname"), c.PostForm("email"), c.PostForm("phone"), c.PostForm("formations"), c.PostForm("experiences"), c.PostForm("competences"), offerId)
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

func SearchCandidatByEmail(c *gin.Context) {
	result, mess := database.SearchCandidatByEmail(c.Query("email"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}

func LoadSomeCandidat(c *gin.Context) {
	result, mess := database.LoadSomeCandidat(c.Query("limit"), c.Query("offset"))
	c.JSON(200, gin.H{
		"message": mess,
		"result":  result,
	})
}
