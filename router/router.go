package router

import "github.com/gin-gonic/gin"

func Initializa() {

	// initialize router
	router := gin.Default()

	// initialize routes
	initializeRoutes(router)

	// linsten and serve on 0.0.0.0:8080
	router.Run(":8080")
}
