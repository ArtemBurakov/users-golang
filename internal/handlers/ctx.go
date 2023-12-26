package handlers

import (
	"github.com/gin-gonic/gin"
	"users/internal/data"
)

// usersQKey is the key used to store the UserQ instance in the Gin context.
const usersQKey = "UsersQ"

// CtxUsersQ is a middleware that sets the UserQ instance in the Gin context.
// It allows you to store a data.UserQ implementation in the context so that
// it can be easily accessed by other handlers in the request lifecycle.
func CtxUsersQ(entry data.UserQ) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(usersQKey, entry)
		c.Next()
	}
}

// UsersQ retrieves the UserQ instance from the Gin context.
// It assumes that the UserQ instance has been set using CtxUsersQ middleware.
// It then calls the New() method on the UserQ instance to get a new instance.
func UsersQ(c *gin.Context) data.UserQ {
	return c.MustGet(usersQKey).(data.UserQ).New()
}
