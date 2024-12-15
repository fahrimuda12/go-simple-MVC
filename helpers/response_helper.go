package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(data interface{}, message string) gin.H {
	return gin.H{"status": http.StatusOK, "data": data, "message":message,"error": false}
}

func ErrorResponse(message string, code string) gin.H {
	return gin.H{"error": true, "message": message, "status": code}
}

func ServerErrorResponse(error string) gin.H {
	return gin.H{"error": error, "message": "Server Error", "status": http.StatusInternalServerError}
}