package main

import (
	api "github.com/Eznopot/RM_api/src/API"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	logger.Info("Server launched")
	router := gin.New()
	router.SetTrustedProxies(nil)
	api.ApiInit(router)
	router.Run("0.0.0.0:8080")
}