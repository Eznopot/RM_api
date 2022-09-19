package Middleware

import (
	"net/http"
	"github.com/Eznopot/RM_api/src/Database"
	"github.com/gin-gonic/gin"
)

func CheckSession(c *gin.Context) {
	token, exist := c.Request.Header["Token"];
	if (!exist) {
		c.AbortWithStatus(http.StatusUnauthorized);
	} else {
		_, res := database.CheckSession(token[0]);
		if (!res) {
			c.AbortWithStatus(http.StatusUnauthorized);
		} else {
			c.Next();
		}
	}
}

func CheckRight(c *gin.Context) {
	token, exist := c.Request.Header["Token"];
	if (!exist) {
		c.AbortWithStatus(http.StatusUnauthorized);
	} else {
		user_id, res := database.CheckSession(token[0]);
		if (res) {
			role, exist := database.CheckRightIsAdmin(user_id)
			if (role == 3 && exist) {
				c.Next();
			} else {
				c.AbortWithStatus(http.StatusUnauthorized);
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized);
		}
	}
}