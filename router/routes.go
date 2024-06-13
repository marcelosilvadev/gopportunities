package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelosilvadev/gopportunities.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
