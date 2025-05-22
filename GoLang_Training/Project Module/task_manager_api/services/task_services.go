package services

import (
	"task_manager_api/model"
	"task_manager_api/repository"

	"github.com/google/uuid"
)

type TaskService struct {
	taskrep repository.TaskManagerRepository
}

func NewTaskService(t repository.TaskManagerRepository) *TaskService {
	return &TaskService{taskrep: t}
}

func (t *TaskService) CreateTask(task *model.Task) error {
	task.ID = uuid.New()
	_, err := t.taskrep.CreateTask(task)

	return err
}

func (t *TaskService) UpdateTask(task *model.Task) (model.Task, error) {
	// task.ID = id
	updatedTask, err := t.taskrep.UpdateTask(task)
	return updatedTask, err
}

func (t *TaskService) DeleteTask(id uuid.UUID) error {
	err := t.taskrep.DeleteTask(id)
	return err
}

func (t *TaskService) GetTaskByID(id uuid.UUID) (model.Task, error) {
	task, err := t.taskrep.GetTaskByID(id)
	return task, err
}

func (t *TaskService) GetTaskList() ([]model.Task, error) {
	tasks, err := t.taskrep.GetTaskList()
	return tasks, err
}
