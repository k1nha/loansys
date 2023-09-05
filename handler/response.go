package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ResponseError(gContext *gin.Context, httpCode int, message string) {
	gContext.Header("Content-Type", "application/json")
	gContext.JSON(httpCode, gin.H{
		"Message":    message,
		"StatusCode": httpCode,
	})
}

func ResponseSuccess(gContext *gin.Context, op string, data interface{}) {
	gContext.Header("Content-Type", "application/json")
	gContext.JSON(200, gin.H{
		"Message": fmt.Sprintf("Operation from handler: %s successfully", op),
		"data":    data,
	})

}
