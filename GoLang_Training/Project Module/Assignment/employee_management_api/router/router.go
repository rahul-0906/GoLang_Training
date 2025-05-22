package router

import (
	"employee_management_api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/employees", handler.CreateEmployees)
	router.GET("/employees", handler.GetEmployees)
	router.GET("/employees/:id", handler.GetEmployee)
	router.PUT("employees/:id", handler.UpdateEmployee)
	router.DELETE("employees/:id", handler.DeleteEmployee)
	return router
}
