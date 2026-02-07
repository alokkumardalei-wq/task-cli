package main

import (
	"fmt"
	"os"
	"strconv"
)

const taskFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	tasks, err := LoadTasks(taskFile)
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task description")
			return
		}
		description := os.Args[2]
		tasks = AddTask(tasks, description)
		if err := SaveTasks(taskFile, tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", tasks[len(tasks)-1].ID)

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Printf("%-4s %-20s %-10s %s\n", "ID", "Status", "Created", "Description")
		fmt.Println("---------------------------------------------------------")
		for _, task := range tasks {
			status := " "
			if task.Completed {
				status = "✓"
			}
			fmt.Printf("%-4d [%s] %-20s %s\n", task.ID, status, task.CreatedAt.Format("2006-01-02"), task.Description)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID: %v\n", err)
			return
		}
		tasks, err = CompleteTask(tasks, id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if err := SaveTasks(taskFile, tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}
		fmt.Printf("Task %d marked as completed.\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID: %v\n", err)
			return
		}
		tasks, err = DeleteTask(tasks, id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if err := SaveTasks(taskFile, tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}
		fmt.Printf("Task %d deleted.\n", id)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: task-cli <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>  Add a new task")
	fmt.Println("  list               List all tasks")
	fmt.Println("  complete <id>      Mark a task as completed")
	fmt.Println("  delete <id>        Delete a task")
}
