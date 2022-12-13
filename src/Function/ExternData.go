package function

import (
	"strconv"

	"github.com/Eznopot/RM_api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetDayOff(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	var res map[string]interface{}
	if err != nil {
		res = utils.GetHollidays(-1)
	} else {
		res = utils.GetHollidays(year)
	}
	c.JSON(200, gin.H{
		"message": res,
		"result":  true,
	})
}
