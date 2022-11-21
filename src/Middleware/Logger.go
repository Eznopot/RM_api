package Middleware

import (
	"strconv"

	database "github.com/Eznopot/RM_api/src/Database"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func CustomLogger(c *gin.Context) {
	token, exist := c.Request.Header["Token"]
	if !exist {
		logger.Info(c.Request.Method, c.Request.URL.Path);
	} else {
		_, res := database.CheckSession(token[0])
		logger.Info(c.Request.Method, c.Request.URL.Path, "by:", strconv.Itoa(res));
	}
	c.Next()
}
