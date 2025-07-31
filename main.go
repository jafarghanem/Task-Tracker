package main

import "fmt"

var Task struct {
	ID          uint
	Description string
	Status      string // "todo", "in-progress", "done"
	CreatedAt   string // timestamp
	UpdatedAt   string // timestamp
}

func main() {

	fmt.Println("This is a simpple task-tracker./nyou can add tasks, track their status, and mark them as the done")
	fmt.Println("To add a task, use the command: add <task_name>")
	fmt.Println("To view all tasks, use the command: list")
	fmt.Println("To Update a task, use the command: update <task_id> <new_info>")
	fmt.Println("To Delete a task, use the command: delete <task_id>")
	fmt.Println("To mark a task as in-progress, use the command: mark-in-progress <task_id>")
	fmt.Println("To mark a task as done, use the command: mark-done <task_id>")
	fmt.Println("You Can List tasks by their status using: list <done|in-progress|todo>")
	fmt.Println("Add The task here and it will be saved in the task.json file")
	fmt.Println("Start From Here : ")
	fmt.Println("--------------------------------------------------")
	// fmt.Scanln()
	
}
