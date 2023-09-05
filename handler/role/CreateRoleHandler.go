package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucascmpus/lend/handler"
	"github.com/lucascmpus/lend/schemas"
)

func CreateRoleHandler(c *gin.Context) {
	logger, db := handler.InitializeHandler()

	request := handler.CreateRoleRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		handler.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	role := schemas.Role{
		Name: request.Name,
	}

	if err := db.Create(role).Error; err != nil {
		logger.ErrorF("Error creating {ROLE} %v", err.Error())
		handler.ResponseError(c, http.StatusInternalServerError, "error on creating {ROLE}")
		return
	}

	handler.ResponseSuccess(c, "create role", role)
}
