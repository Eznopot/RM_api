package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Eznopot/RM_api/src/API"
)

func main() {
	fmt.Println("server launched")
	router := gin.New()
	api.ApiInit(router)
	router.Run("0.0.0.0:8080")
}

