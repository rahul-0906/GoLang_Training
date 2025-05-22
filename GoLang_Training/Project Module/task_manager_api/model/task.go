package model

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a task with its properties.
type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TaskAssignment struct {
	ID         string    `json:"id"`
	TaskID     string    `json:"task_id"`
	UserID     string    `json:"user_id"`
	AssignedAt time.Time `json:"assigned_at"`
}

// Mock data for tasks
// var Tasks = []Task{
// 	{ID: "ee3b8b9f-debf-4cd5-9551-03b92dc4b660", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
// 	{ID: "0c5a609d-357e-4684-a277-e95dc9c48002", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
// 	{ID: "72607711-7708-4735-a63b-fdc442228801", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
// }
