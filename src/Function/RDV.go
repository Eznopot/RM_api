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
	candidatId, _ := strconv.Atoi(c.PostForm("candidatId"))
	res, mess := database.AddRDVEvent(c.Request.Header["Token"][0], candidatId, c.PostForm("date"))

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
	id, _ := strconv.Atoi(c.PostForm("id"))
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