package repository

import (
	"fmt"
	"task_manager_api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskManagerRepository struct {
	DB *gorm.DB
}

type TaskRepository interface {
	CreateTask(task *model.Task) (model.Task, error)
	UpdateTask(task *model.Task) (model.Task, error)
	DeleteTask(id uuid.UUID) (model.Task, error)
	GetTaskByID(id uuid.UUID) (model.Task, error)
	GetTaskList() ([]model.Task, error)
}

func NewTaskManagerRepository(db *gorm.DB) *TaskManagerRepository {
	return &TaskManagerRepository{DB: db}
}

func (tm *TaskManagerRepository) CreateTask(task *model.Task) (model.Task, error) {
	result := tm.DB.Create(&task)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var createdTask model.Task
	result.Scan(&createdTask)

	return createdTask, result.Error
}

func (tm *TaskManagerRepository) UpdateTask(task *model.Task) (model.Task, error) {

	result := tm.DB.Save(task)
	var updatedTask *model.Task
	result.Scan(&updatedTask)

	return *updatedTask, result.Error
}

func (tm *TaskManagerRepository) DeleteTask(id uuid.UUID) error {

	return tm.DB.Delete(&model.Task{}, "id = ?", id).Error

}

func (tm *TaskManagerRepository) GetTaskByID(id uuid.UUID) (model.Task, error) {

	var task model.Task
	err := tm.DB.First(&task, "id = ?", id).Error
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (tm *TaskManagerRepository) GetTaskList() ([]model.Task, error) {

	var tasks []model.Task
	err := tm.DB.Find(&tasks).Error

	return tasks, err
}
