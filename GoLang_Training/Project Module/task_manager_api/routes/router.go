package router

import (
	"net/http"
	"task_manager_api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(th *handler.TaskHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server is up",
		})
	})

	apiRouter := router.Group("/api")
	taskRouter := apiRouter.Group("/task")
	taskRouter.GET("/", th.GetTaskList)
	taskRouter.POST("/create", th.CreateTask)
	taskRouter.GET("/:taskID", th.GetTaskByID)
	taskRouter.PATCH("/:taskID", th.UpdateTask)
	taskRouter.DELETE("/:taskID", th.DeleteTask)

	userRouter := apiRouter.Group("user")
	userRouter.GET("")
	userRouter.POST("/:userID")
	userRouter.GET("/:userID")
	userRouter.PATCH("/:userID")
	userRouter.DELETE("/:userID")

	return router
}
