package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddOffer(c *gin.Context) {
	res, mess := database.AddOffer(c.Request.Header["Token"][0], c.PostForm("title"), c.PostForm("description"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func ModifyOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.ModifyOffer(id, c.PostForm("title"), c.PostForm("description"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.DeleteOffer(id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func GetOffers(c *gin.Context) {
	res, mess := database.GetOffers()

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}
