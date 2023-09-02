package role

import (
	"github.com/gin-gonic/gin"
	"github.com/lucascmpus/lend/handler"
)

func CreateRoleHandler(c *gin.Context) {
	logger, db := handler.InitializeHandler()

	request := handler.CreateRoleRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("")
	}

	if err := db.Create(&request).Error; err != nil {
		logger.ErrorF("Error creating {ROLE} %v", err.Error())
		return
	}

	// c.JSON(http.StatusOK, )

}
