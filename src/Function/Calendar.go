package function

import (
	"os"
	strconv "strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/Eznopot/RM_api/src/utils"
	"github.com/gin-gonic/gin"
)

func AddCalendarEvent(c *gin.Context) {
	value, err := strconv.ParseFloat(c.PostForm("value"), 32)
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

func GetCalendarEvents(c *gin.Context) {
	monthNbr, _ := strconv.Atoi(c.Query("month"))
	res, mess := database.GetCalendarEvents(c.Request.Header["Token"][0], monthNbr)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func AutoPresenceCalendarEvents(c *gin.Context) {
	monthNbr, _ := strconv.Atoi(c.PostForm("month"))
	res, mess := database.AutoPresenceCalendarEvents(c.Request.Header["Token"][0], monthNbr)

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func ModifyCalendarEvent(c *gin.Context) {
	value, err := strconv.ParseFloat(c.PostForm("value"), 32)
	id, err2 := strconv.Atoi(c.PostForm("id"))
	if err != nil || err2 != nil {
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.ModifyCalendarEvent(c.Request.Header["Token"][0], id, c.PostForm("date"), c.PostForm("eventType"), c.PostForm("comment"), value, c.PostForm("otherValue"), c.PostForm("backupName"), c.PostForm("absenceType"))

	c.JSON(200, gin.H{
		"message": mess,
		"result":  res,
	})
}

func DeleteCalendarEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		println(err.Error())
		c.JSON(200, gin.H{
			"message": "value is not an int",
			"result":  false,
		})
		return
	}
	res, mess := database.DeleteCalendarEvent(c.Request.Header["Token"][0], id)

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

func GetOwnCram(c *gin.Context) {
	monthNbr, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{
			"message": "bad parameters",
			"result":  false,
		})
		return
	}
	complete, res := database.GetCalendarEvents(c.Request.Header["Token"][0], monthNbr)
	if !complete {
		logger.Error(err.Error())
		c.JSON(500, gin.H{
			"message": "error retrieving data",
			"result":  false,
		})
		return
	}
	complete, user := database.GetInfo(c.Request.Header["Token"][0])
	filename := utils.CreateExcelFileAndSaveIt(res, user, monthNbr)
	c.FileAttachment(filename, filename)
	os.Remove(filename)
}

// dont work yet
func GetAllCram(c *gin.Context) {
	monthNbr, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{
			"message": "bad parameters",
			"result":  false,
		})
		return
	}

	complete, users := database.GetAllUserWithAllData()

	if !complete {
		logger.Error("can't get all user")
		c.JSON(500, gin.H{
			"message": "cant get all user",
			"result":  false,
		})
		return
	}

	listFile := []string{}

	for _, user := range users {
		complete, res := database.GetCalendarEventsByUserId(user.Id, monthNbr)
		if !complete {
			logger.Error(err.Error())
			c.JSON(500, gin.H{
				"message": "error retrieving data",
				"result":  false,
			})
			return
		}
		filename := utils.CreateExcelFileAndSaveIt(res, user, monthNbr)
		println("dl file:", filename)
		listFile = append(listFile, filename)
	}

	zipFile, complete := utils.ZipFiles(monthNbr, listFile)

	for _, filename := range listFile {
		os.Remove(filename)
	}
	c.File(zipFile)
	os.Remove(zipFile)
}
