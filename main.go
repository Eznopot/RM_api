package main

import (
	"fmt"

	api "github.com/Eznopot/RM_api/src/API"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("server launched")
	
	router := gin.New()
	router.SetTrustedProxies(nil)
	api.ApiInit(router)
	router.Run("0.0.0.0:8080")
}
