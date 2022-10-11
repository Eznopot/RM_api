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

/*
List des Idée/changement a faire:
- Mettre dabords present partout Calendar (Réflechir si chrone qui s'execute tout les 1er du mois pourrait marcher)
- Faire un sorte que toute les personne au siege voient si qqn est en congés (peut etre rajouter un role en plsu et faire une page spécifiquement pour ca)
*/