package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddOffer(c *gin.Context) {
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not a float",
			"result":  false,
		})
		return
	}
	res, mess := database.AddOffer(c.Request.Header["Token"][0], c.PostForm("title"), c.PostForm("description"), price)

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
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not a float",
			"result":  false,
		})
		return
	}
	res, mess := database.ModifyOffer(id, c.PostForm("title"), c.PostForm("description"), price)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

func GetSomeOffers(c *gin.Context) {
	res, mess := database.LoadSomeOffers(c.Query("limit"), c.Query("offset"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}
