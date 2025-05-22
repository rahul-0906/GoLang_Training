package handler

import (
	"fmt"
	"net/http"
	"task_manager_api/model"
	"task_manager_api/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	Service *services.TaskService
}

func NewTaskHandler(ts *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: ts}
}

func (th *TaskHandler) CreateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := th.Service.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, task)
	return
}

func (th *TaskHandler) UpdateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	taskID := ctx.Param("taskID")
	task.ID = uuid.MustParse(taskID)

	updatedTask, err := th.Service.UpdateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
	return
}

func (th *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("taskID")

	err := th.Service.DeleteTask(uuid.MustParse(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
	return
}

func (th *TaskHandler) GetTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("taskID")

	task, err := th.Service.GetTaskByID(uuid.MustParse(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("task", task)
	ctx.JSON(http.StatusOK, task)
	return
}

func (th *TaskHandler) GetTaskList(ctx *gin.Context) {
	tasks, err := th.Service.GetTaskList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("tasks", tasks)
	ctx.JSON(http.StatusOK, tasks)
	return
}
