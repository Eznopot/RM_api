package function

import (
	strconv "strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func GetRDVEvent(c *gin.Context) {
	monthNbr, _ := strconv.Atoi(c.PostForm("month"))
	res, mess := database.GetRDVEvent(monthNbr)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func AddRDVEvent(c *gin.Context) {
	res, mess := database.AddRDVEvent(c.Request.Header["Token"][0], c.PostForm("email"), c.PostForm("firstName"), c.PostForm("lastName"), c.PostForm("date"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func ModifyRDVEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	res, mess := database.ModifyRDVEvent(c.Request.Header["Token"][0], id, c.PostForm("date"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteRDVEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, mess := database.DeleteRDVEvent(c.Request.Header["Token"][0], id)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func InsertAppreciation(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	res, mess := database.InsertAppreciation(c.Request.Header["Token"][0], id, c.PostForm("appreciation"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

/*



*/