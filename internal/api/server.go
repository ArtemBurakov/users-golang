package api

import (
	"github.com/gin-gonic/gin"
	"users/internal/data/memoryStore"
	"users/internal/handlers"
)

type Server struct {
	router *gin.Engine
	store  *memoryStore.MemoryStore
}

// CORSMiddleware is a middleware for handling Cross-Origin Resource Sharing (CORS) headers.
// Only for local run.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// NewServer creates a new instance of the API server with the given data store and configuration.
func NewServer(store *memoryStore.MemoryStore) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.Use(CORSMiddleware(), handlers.CtxUsersQ(memoryStore.NewUserQ(server.store)))

	api := router.Group("/api")
	{
		user := api.Group("/users")
		user.GET("/", handlers.GetUsers)
		user.POST("/", handlers.AddUser)
		user.GET("/:id", handlers.GetUser)
		user.PUT("/:id", handlers.UpdateUser)
		user.DELETE("/:id", handlers.DeleteUser)
	}

	server.router = router
	return server
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}
