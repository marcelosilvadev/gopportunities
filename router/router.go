package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o router utilizando a configurações Default do Gin
	r := gin.Default()

	//Definindo uma Rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Roda a Api
	r.Run()
}
