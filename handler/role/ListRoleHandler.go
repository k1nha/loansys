package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucascmpus/lend/config"
	"github.com/lucascmpus/lend/schemas"
)

func ListRoleHandler(c *gin.Context) {
	db := config.GetDB()

	cargos := db.Find(&schemas.Role{})

	c.JSON(http.StatusOK, cargos)
}
