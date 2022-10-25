package api

import (
	function "github.com/Eznopot/RM_api/src/Function"
	"github.com/Eznopot/RM_api/src/Middleware"
	"github.com/Eznopot/RM_api/src/utils"
	"github.com/gin-gonic/gin"
)

func ApiInit(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	utils.BetterPrint("", "No rights route")
	router.POST("/api/test", function.Test)

	router.POST("/user/login", function.Login)

	router.GET("/calendar/getEventTypes", function.GetEventTypes)
	router.GET("/calendar/getOtherEventTypes", function.GetOtherEventTypes)
	router.GET("/calendar/getAbsenceEventTypes", function.GetAbsenceEventTypes)
	router.DELETE("/user/logout", function.Logout)

	utils.BetterPrint(utils.Green, "User route", true)
	logged := router.Group("/session")
	logged.Use(Middleware.CheckSession)
	{
		logged.GET("/user/getInfo", function.GetInfo)
		logged.GET("/user/getPages", function.GetPages)
		logged.GET("/user/checkSession", function.CheckSession)
		logged.POST("/user/addCalendarEvent", function.AddCalendarEvent)
		logged.GET("/user/getCalendarEvents", function.GetCalendarEvents)
		logged.POST("/user/autoPresenceCalendarEvents", function.AutoPresenceCalendarEvents)

		logged.POST("/user/modifyCalendarEvent", function.ModifyCalendarEvent)
		logged.POST("/user/deleteCalendarEvent", function.DeleteCalendarEvent)

		logged.GET("/user/getHollidayRequest", function.GetHollidayRequest)
		logged.POST("/user/addHollidayRequest", function.AddHollidayRequest)
		logged.POST("/user/deleteHollidayRequest", function.DeleteHollidayRequest)
	}

	utils.BetterPrint(utils.Yellow, "Manager route", true)
	manager := router.Group("/sessionManager")
	manager.Use(Middleware.CheckManager)
	{
		manager.GET("/holliday/getAllHollidayRequest", function.GetAllHollidayRequest)
	}

	utils.BetterPrint(utils.Red, "Admin route", true)
	admin := router.Group("/sessionPlus")
	admin.Use(Middleware.CheckAdmin)
	{
		admin.GET("/user/getAllUser", function.GetAllUser)
		admin.POST("/user/updateRole", function.UpdateRole)
		admin.POST("user/register", function.Register)

		admin.POST("/candidat/add", function.AddCandidat)
		admin.GET("/candidat/search", function.SearchCandidat)
	}
	utils.ResetColor()
}
