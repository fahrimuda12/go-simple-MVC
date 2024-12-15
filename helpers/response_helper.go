package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func successResponse(data interface{}, message string) gin.H {
	return gin.H{"status": http.StatusOK, "data": data, "message":message,"error": false}
}

func errorResponse(message string, code string) gin.H {
	return gin.H{"error": true, "message": message, "status": code}
}

func serverErrorResponse(error string) gin.H {
	return gin.H{"error": error, "message": "Server Error", "status": http.StatusInternalServerError}
}