package Middleware

import (
	"fmt"

	database "github.com/Eznopot/RM_api/src/Database"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func CustomLogger(c *gin.Context) {
	token, exist := c.Request.Header["Token"]
	if !exist {
		logger.Info(fmt.Sprintf("[%s] %s - Status code: %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status()))
	} else {
		_, res := database.CheckSession(token[0])
		logger.Info(fmt.Sprintf("[%s] %s - Status code: %d - by %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), res))
	}
	c.Next()
}
