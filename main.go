package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`    // "todo", "in-progress", "done"
	CreatedAt   string `json:"createdat"` // timestamp
	UpdatedAt   string `json:"updatedat"` // timestamp
}

func main() {
	fmt.Println("This is a simpple task-tracker.")
	fmt.Println("you can add tasks, track their status, and mark them as the done")
	fmt.Println("To add a task, use the command: add \"task_name\"")
	fmt.Println("To view all tasks, use the command: list")
	fmt.Println("To Update a task, use the command: update <task_id> <new_info>")
	fmt.Println("To Delete a task, use the command: delete <task_id>")
	fmt.Println("To mark a task as in-progress, use the command: mark-in-progress <task_id>")
	fmt.Println("To mark a task as done, use the command: mark-done <task_id>")
	fmt.Println("You Can List tasks by their status using: list <done|in-progress|todo>")
	fmt.Println("Add The task here and it will be saved in the task.json file")
	fmt.Println("Start From Here : ")
	fmt.Println("--------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
		}
		input = strings.TrimSpace(input) // Remove any trailing newline or spaces

		Task := Task{}
		usrcommand, args := parseInput(input)
		switch usrcommand {
		case "add":
			Task.Description = args[0]
			Task.CreatedAt = NowTimeString()
			Task.Status = "todo"
			AddTaskToJson(&Task)
		case "update":
			uid, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid number.")
			}
			Task.Description = args[1]
			UpdateTasktoJson(uint(uid), &Task)
		case "delete":
			uid, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid number.")
			}
			DeleteTaskFromJson(uint(uid))
		case "list":
			if len(args) == 0 {
				ListTasksFromJson("") // list all tasks
			} else {
				ListTasksFromJson(args[0]) // list tasks by status
			}
		case "mark-in-progress":
			uid, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid number.")
			}
			MarkTaskInProgressJson(uint(uid))
		case "mark-done":
			uid, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid number.")
			}
			MarkTaskDoneJson(uint(uid))
		case "exit":
			fmt.Println("Exiting the task tracker. Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}

func parseInput(input string) (string, []string) {
	var parts []string
	var current strings.Builder
	inQuotes := false

	for _, r := range input {
		switch {
		case r == '"':
			inQuotes = !inQuotes
		case r == ' ' && !inQuotes:
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

func NowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func readTasksFromJSON() ([]Task, error) {
	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	info, _ := file.Stat()
	if info.Size() == 0 {
		return tasks, nil // Return empty slice if file is empty
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func writeTasksToJSON(tasks []Task) error {
	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func AddTaskToJson(t *Task) error {
	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read existing tasks
	var tasks []Task
	stat, _ := file.Stat()
	if stat.Size() > 0 {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&tasks); err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}
	}

	// Assign next ID
	var nextID uint = 1
	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
	t.ID = nextID
	tasks = append(tasks, *t)

	// Reset file content before writing
	file.Truncate(0)
	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error encoding tasks to JSON: %v", err)
	}

	fmt.Printf("Task added successfully (ID: %v).\n", t.ID)
	return nil
}

func UpdateTasktoJson(id uint, updated *Task) error {
	tasks, err := readTasksFromJSON()
	if err != nil {
		return err
	}

	updatedAt := NowTimeString()
	updated.UpdatedAt = updatedAt

	updatedFlag := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = updated.Description
			tasks[i].Status = updated.Status
			tasks[i].UpdatedAt = updatedAt
			updatedFlag = true
			fmt.Printf("Task with ID %d updated successfully.\n", id)
			break
		}
	}
	if !updatedFlag {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return writeTasksToJSON(tasks)
}

func DeleteTaskFromJson(id uint) {
	tasks, err := readTasksFromJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newTasks := tasks[:0]
	found := false
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}

	if err := writeTasksToJSON(newTasks); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task deleted successfully.")
}

func ListTasksFromJson(filterStatus string) {
	file, err := os.Open("task.json")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Printf("error decoding JSON: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		if filterStatus == "" || task.Status == filterStatus {
			fmt.Printf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func MarkTaskInProgressJson(id uint) {
	tasks, err := readTasksFromJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = NowTimeString()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}

	if err := writeTasksToJSON(tasks); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task marked as in-progress.")
}

func MarkTaskDoneJson(id uint) {
	tasks, err := readTasksFromJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = NowTimeString()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}

	if err := writeTasksToJSON(tasks); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task marked as done.")
}
