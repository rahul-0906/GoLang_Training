package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}
