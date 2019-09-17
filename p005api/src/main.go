package main

import (
	"github.com/gin-gonic/gin"
	handler "translate/P005api/src/handler/p010user/v1"
)

func main() {
	r:= gin.Default()


	api := r.Group("/api")
	{
		// Health Check
		api.GET("/health", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"message": "Alive",
			})
		})

		v1 := api.Group("/v1")
		{
			v1.POST("/user/login", handler.LoginHandler())
		}
	}
	r.Run(":8080")

}