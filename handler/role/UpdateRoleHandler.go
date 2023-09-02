package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Updates Cargo",
	})
}
