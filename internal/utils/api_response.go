package utils

import "github.com/gin-gonic/gin"

func Null() interface{} {
	return nil
}

func ApiResponse[T any](status string, data T) gin.H {
	return gin.H{"status": status, "data": data}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
