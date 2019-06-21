package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router:= gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/register", func(context *gin.Context) {

		})

		v1.GET("/show", func(context *gin.Context) {
			router.Use()
		})
	}



	router.Run(":3001")


}

