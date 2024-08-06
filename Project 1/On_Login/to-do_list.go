package On_Login

import (
	"FileHandling/utils"
	"encoding/json"
	"fmt"
	"os"
)

// Task struct for a single line task
type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// UserTasks struct to hold tasks for each user
type UserTasks struct {
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}

var userTasks []UserTasks

const tasksFilename = "tasks.json"

// Load tasks from JSON file
func loadTasks() error {
	file, err := os.ReadFile(tasksFilename)
	if err != nil {
		if os.IsNotExist(err) {
			userTasks = []UserTasks{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &userTasks)
}

// Save tasks to JSON file
func saveTasks() error {
	file, err := json.MarshalIndent(userTasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFilename, file, 0644)
}

// Add a new task
func addTask() {
	title := utils.ReadInput("Enter task : ")

	newTask := Task{
		Title:     title,
		Completed: false,
	}

	var userFound bool
	for i, user := range userTasks {
		if user.Username == ActiveUser.Username {
			userTasks[i].Tasks = append(userTasks[i].Tasks, newTask)
			userFound = true
			break
		}
	}

	if !userFound {
		userTasks = append(userTasks, UserTasks{
			Username: ActiveUser.Username,
			Tasks:    []Task{newTask},
		})
	}

	if err := saveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	} else {
		fmt.Println("Task added successfully!")
	}
}

// Mark a task as complete
func markTaskComplete() {
	taskTitle := utils.ReadInput("Enter the title of the task to mark as complete: ")

	var userFound bool
	for i, user := range userTasks {
		if user.Username == ActiveUser.Username {
			for j, task := range user.Tasks {
				if task.Title == taskTitle && !task.Completed {
					userTasks[i].Tasks[j].Completed = true
					userFound = true
					break
				}
			}
			break
		}
	}

	if !userFound {
		fmt.Println("Task not found or already completed.")
	} else {
		if err := saveTasks(); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		} else {
			fmt.Println("Task marked as complete successfully!")
		}
	}
}

// View completed tasks
func viewCompletedTasks() {
	fmt.Println("Completed Tasks:")
	for _, user := range userTasks {
		if user.Username == ActiveUser.Username {
			for _, task := range user.Tasks {
				if task.Completed {
					fmt.Printf("Title: %s\n", task.Title)
				}
			}
			return
		}
	}
	fmt.Println("No completed tasks found.")
}

// View uncompleted tasks
func viewUncompletedTasks() {
	fmt.Println("Uncompleted Tasks:")
	for _, user := range userTasks {
		if user.Username == ActiveUser.Username {
			for _, task := range user.Tasks {
				if !task.Completed {
					fmt.Printf("Title: %s\n", task.Title)
				}
			}
			return
		}
	}
	fmt.Println("No uncompleted tasks found.")
}

// Task management section
func TaskManagementSection() {
	for {
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;34m              TO-DO LIST          \033[0m")            // Blue
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")
		fmt.Println()
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Complete")
		fmt.Println("3. View Completed Tasks")
		fmt.Println("4. View Uncompleted Tasks")
		fmt.Println("5. Exit")
		fmt.Println()

		var choice int
		fmt.Print("\033[1;34mEnter your choice: \033[0m") // Blue

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\033[1;31mError reading input. Please enter a valid number.\033[0m") // Red
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			markTaskComplete()
		case 3:
			viewCompletedTasks()
		case 4:
			viewUncompletedTasks()
		case 5:
			return // Exit the function
		default:
			fmt.Println("\033[1;31mInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
