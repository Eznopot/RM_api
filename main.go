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

/*
List des Idée/changement a faire:
- Mettre dabords present partout Calendar (Réflechir si chrone qui s'execute tout les 1er du mois pourrait marcher)
- Faire un sorte que toute les personne au siege voient si qqn est en congés (peut etre rajouter un role en plsu et faire une page spécifiquement pour ca)



_, res := database.GetCalendarEvents("3fcbe062-df02-457f-be15-66f0c805fa92", 11)
	utils.CreateExcelFileAndSaveIt(res, 11)
	return
*/
