package Middleware

import (
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func CustomLogger(c *gin.Context) {
	logger.Info(c.Request.Method, c.Request.URL.Path)
	c.Next()
}
