package api

import (
	"time"

	function "github.com/Eznopot/RM_api/src/Function"
	"github.com/Eznopot/RM_api/src/Middleware"
	"github.com/Eznopot/RM_api/src/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiInit(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(Middleware.CustomLogger)
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	utils.BetterPrint("", "No rights route")
	router.POST("/api/test", function.Test)

	router.POST("/user/login", function.Login)

	router.GET("/getDayOff", function.GetDayOff)

	router.GET("/calendar/getEventTypes", function.GetEventTypes)
	router.GET("/calendar/getOtherEventTypes", function.GetOtherEventTypes)
	router.GET("/calendar/getAbsenceEventTypes", function.GetAbsenceEventTypes)
	router.POST("/candidat/add", function.AddCandidat)
	router.GET("/info/roles", function.GetRoles)
	router.GET("/offer/getAll", function.GetOffers)
	router.GET("/offer/getSome", function.GetSomeOffers)

	utils.BetterPrint(utils.Green, "User route", true)
	logged := router.Group("/session")
	logged.Use(Middleware.CheckSession)
	{
		logged.DELETE("/user/logout", function.Logout)
		
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

		logged.GET("/user/getOwnCRAM", function.GetOwnCram)

		logged.GET("/info/get", function.GetAdminString)
	}

	utils.BetterPrint(utils.Yellow, "Manager route", true)
	manager := router.Group("/sessionManager")
	manager.Use(Middleware.CheckManager)
	{
		manager.GET("/holliday/getAllHollidayRequest", function.GetAllHollidayRequest)
		manager.POST("/holliday/declineHollidayRequest", function.DeclineHollidayRequest)
		manager.POST("/holliday/acceptHollidayRequest", function.AcceptHollidayRequest)
		manager.POST("/holliday/deleteOtherHollidayRequest", function.DeleteOtherHollidayRequest)

		manager.GET("/candidat/search", function.SearchCandidat)
		manager.GET("/candidat/searchByEmail", function.SearchCandidatByEmail)
		manager.GET("/candidat/loadSome", function.LoadSomeCandidat)

		manager.GET("/RDV/getAll", function.GetRDVEvent)
		manager.POST("/RDV/add", function.AddRDVEvent)
		manager.POST("/RDV/modify", function.ModifyRDVEvent)
		manager.DELETE("/RDV/delete/:id", function.DeleteRDVEvent)
		manager.POST("/RDV/saveNote", function.InsertAppreciation)
		manager.POST("/offer/add", function.AddOffer)
		manager.POST("/offer/modify", function.ModifyOffer)
		manager.DELETE("/offer/delete/:id", function.DeleteOffer)
	}

	utils.BetterPrint(utils.Red, "Admin route", true)
	admin := router.Group("/sessionPlus")
	admin.Use(Middleware.CheckAdmin)
	{
		admin.GET("/user/getAllUser", function.GetAllUser)
		admin.POST("/user/updateRole", function.UpdateRole)
		admin.POST("user/register", function.Register)

		admin.POST("/info/add", function.AddAdminString)
		admin.POST("/info/modify", function.ModifyAdminString)
		admin.POST("/info/delete", function.DeleteAdminString)

		admin.GET("/calendar/getAllCRAM", function.GetAllCram)
		admin.GET("/user/getCRAMByEmail", function.GetCrambyEmail)

		admin.GET("/user/cv/get", function.GetUserCv)

	}
	utils.ResetColor()
}
