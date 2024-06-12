package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o router utilizando a configurações Default do Gin
	router := gin.Default()

	//Inicializando Routes
	initializeRoutes(router)
	//Roda a Api
	router.Run()
}
