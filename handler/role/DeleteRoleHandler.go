package role

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucascmpus/lend/config"
	"github.com/lucascmpus/lend/handler"
	"github.com/lucascmpus/lend/schemas"
)

func DeleteRoleHandler(c *gin.Context) {
	db := config.GetDB()
	id := c.Query("id")
	if id == "" {
		handler.ResponseError(c, http.StatusBadRequest, handler.ErrorParamRequired("id", "queryParamether").Error())
	}
	role := schemas.Role{}

	if err := db.First(&role, id).Error; err != nil {
		handler.ResponseError(c, http.StatusFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&role).Error; err != nil {
		handler.ResponseError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting %s", id))
		return
	}

	handler.ResponseSuccess(c, "deleted role", role)
}
