package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Eznopot/RM_api/src/Middleware"
	"github.com/Eznopot/RM_api/src/Function"
)

func ApiInit(router *gin.Engine) {
	router.Use(gin.Recovery());
	router.Use(gin.Logger());

	router.POST("/api/test", function.Test);

	router.POST("user/login", function.Login);
	
	logged := router.Group("/session")
	logged.Use(Middleware.CheckSession)
	{
		logged.DELETE("/user/logout", function.Logout);
		logged.GET("/user/getInfo", function.GetInfo);
		logged.GET("/user/getPages", function.GetPages);
		logged.GET("/user/checkSession", function.CheckSession);
	}
	
	admin := router.Group("/sessionPlus")
	admin.Use(Middleware.CheckRight)
	{
		admin.GET("/user/getAllUser", function.GetAllUser);
		admin.POST("/user/updateRole", function.UpdateRole);
		router.POST("user/register", function.Register);
		
		admin.POST("/candidat/add", function.AddCandidat);
		admin.GET("/candidat/search", function.SearchCandidat);
	}
}