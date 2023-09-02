package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucascmpus/lend/handler/role"
	"github.com/lucascmpus/lend/handler"
)

func initRoutes(r *gin.Engine) {

	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/role", role.CreateRoleHandler)
		v1.GET("/role", role.ListRoleHandler)
		v1.DELETE("/role", role.DeleteRoleHandler)
		v1.PATCH("/role", role.UpdateRoleHandler)
	}
}
