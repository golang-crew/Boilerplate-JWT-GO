package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-crew/Bolierplate-JWT-Auth/handler"
)

func ApplyRoutes(r *gin.RouterGroup) {
	test := r.Group("/ping")
	{
		test.GET("/", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
	}
	todo := r.Group("/todo")
	{
		todo.POST("/todo", handler.CreateTodo)
	}
}
