// utils/response.go
package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"message": message, "data": data})
}
