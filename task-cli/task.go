package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Task represents a component of a to-do list.
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Completed   bool      `json:"completed"`
}

// LoadTasks reads tasks from the specified file.
func LoadTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks writes tasks to the specified file.
func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// AddTask adds a new task to the list.
func AddTask(tasks []Task, description string) []Task {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task := Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}
	return append(tasks, task)
}

// CompleteTask marks a task as completed.
func CompleteTask(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			tasks[i].CompletedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", id)
}

// DeleteTask removes a task from the list.
func DeleteTask(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", id)
}
