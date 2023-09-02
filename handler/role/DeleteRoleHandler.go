package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deletes Cargo",
	})
}
