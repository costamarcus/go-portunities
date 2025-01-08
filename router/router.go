package router

import "github.com/gin-gonic/gin"

func Initializa() {

	// inicializa a router utilizando as configurações default do gin
	router := gin.Default()

	// define uma rota
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// linsten and serve on 0.0.0.0:8080
	router.Run(":8080")
}
