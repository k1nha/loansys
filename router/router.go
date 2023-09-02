package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	r := gin.Default()

	initRoutes(r)

	r.Run(":8000")
}
