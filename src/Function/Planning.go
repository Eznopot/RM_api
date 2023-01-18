package function

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func GetPlanning(c *gin.Context) {
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{
			"message": "bad parameters",
			"result":  false,
		})
		return

	}
	res, mess := database.GetPlanning(month)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func AddPlanning(c *gin.Context) {
	res, mess := database.AddPlanning(c.PostForm("email"), c.PostForm("clientName"), c.PostForm("dateStart"), c.PostForm("dateEnd"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeletePlanning(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, mess := database.DeletePlanning(id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}
