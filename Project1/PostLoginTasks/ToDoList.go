package PostLoginTasks

import (
	"FileHandling/Config"
	"FileHandling/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task struct for a single line task with an ID
type Task struct {
	ID        int    `json:"id"`
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
	file, err := os.ReadFile(Config.TaskFile)
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
	return os.WriteFile(Config.TaskFile, file, 0644)
}

// Generate a unique ID for a new task
func generateTaskID(user UserTasks) int {
	var maxID int
	for _, task := range user.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

// Add a new task

func addTask() {
	title := utils.ReadInput("Enter task: ")

	var newTask Task

	// Find the user and generate a unique ID for the new task
	for i, user := range userTasks {
		if user.Username == ActiveUser.Username {
			newTask = Task{
				ID:        generateTaskID(user),
				Title:     title,
				Completed: false,
			}
			userTasks[i].Tasks = append(userTasks[i].Tasks, newTask)
			break
		}
	}

	// If user not found, create a new entry
	if newTask.ID == 0 {
		newTask = Task{
			ID:        1,
			Title:     title,
			Completed: false,
		}
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

// Mark a task as complete using the task ID
func markTaskComplete() {
	taskIDStr := utils.ReadInput("Enter the ID of the task to mark as complete: ")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	var taskFound bool
	for i, user := range userTasks {
		if user.Username == ActiveUser.Username {
			for j, task := range user.Tasks {
				if task.ID == taskID && !task.Completed {
					userTasks[i].Tasks[j].Completed = true
					taskFound = true
					break
				}
			}
			break
		}
	}

	if !taskFound {
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
					fmt.Printf("%d :  %s\n", task.ID, task.Title)
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
					fmt.Printf("%d : %s\n", task.ID, task.Title)
				}
			}
			return
		}
	}
	fmt.Println("No uncompleted tasks found.")
}

// Task management section

func TaskManagementSection() {
	// Load tasks before displaying the task management menu
	if err := loadTasks(); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

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
