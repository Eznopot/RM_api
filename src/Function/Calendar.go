package function

import (
	strconv "strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func AddCalendarEvent(c *gin.Context) {
	value, err := strconv.ParseFloat(c.PostForm("value"), 2)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "value is not a float",
			"result":  false,
		})
	}
	res, mess := database.AddCalendarEvent(c.Request.Header["Token"][0], c.PostForm("date"), c.PostForm("eventType"), c.PostForm("comment"), value, c.PostForm("otherValue"), c.PostForm("backupName"), c.PostForm("absenceType"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func GetEventTypes(c *gin.Context) {
	res, mess := database.GetEventTypes()
	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}


func GetOtherEventTypes(c *gin.Context) {
	res, mess := database.GetOtherEventTypes()

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}


func GetAbsenceEventTypes(c *gin.Context) {
	res, mess := database.GetAbsenceEventTypes()

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}
