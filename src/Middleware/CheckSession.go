package Middleware

import (
	"net/http"

	database "github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func CheckSession(c *gin.Context) {
	token, exist := c.Request.Header["Token"]
	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		res, _ := database.CheckSession(token[0])
		if !res {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}

func CheckAdmin(c *gin.Context) {
	token, exist := c.Request.Header["Token"]
	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		res, user_id := database.CheckSession(token[0])
		if res {
			exist, role := database.CheckRightIsAdmin(user_id)
			if role >= 2 && exist {
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func CheckManager(c *gin.Context) {
	token, exist := c.Request.Header["Token"]
	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		res, user_id := database.CheckSession(token[0])
		if res {
			exist, role := database.CheckRightIsAdmin(user_id)
			if role >= 1 && exist {
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
