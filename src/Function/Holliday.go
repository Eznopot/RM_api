package function

import (
	strconv "strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddHollidayRequest(c *gin.Context) {
	res, mess := database.AddHollidayRequest(c.Request.Header["Token"][0], c.PostForm("dateStart"), c.PostForm("dateEnd"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func 	AcceptHollidayRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.AcceptHollidayRequest(c.Request.Header["Token"][0], id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeclineHollidayRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.DeclineHollidayRequest(c.Request.Header["Token"][0], id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func ModifyHollidayRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.ModifyHollidayRequest(c.Request.Header["Token"][0], id, c.PostForm("dateStart"), c.PostForm("dateEnd"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteHollidayRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.DeleteHollidayRequest(c.Request.Header["Token"][0], id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteOtherHollidayRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.DeleteOtherHollidayRequest(id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func GetHollidayRequest(c *gin.Context) {
	res, mess := database.GetHollidayRequest(c.Request.Header["Token"][0])

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func GetAllHollidayRequest(c * gin.Context) {
	monthNbr, _ := strconv.Atoi(c.PostForm("month"));
	res, mess := database.GetAllHollidayRequest(c.Request.Header["Token"][0], monthNbr);

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}
