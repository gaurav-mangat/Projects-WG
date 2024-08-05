package On_Login

import (
	a "FileHandling/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Define the structure for a to-do item list

type TodoItem struct {
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

// Define a global variable for to-do list
var todoList []TodoItem

// Load to-do list from JSON file
func loadTodoList(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, no to-do items to load
			todoList = []TodoItem{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &todoList)
}

// Save to-do list to JSON file
func saveTodoList(filename string) error {
	file, err := json.MarshalIndent(todoList, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// Add a new to-do item

func AddTodoItem() {
	const filename = "todo_list.json"

	// Load to-do list from the file
	if err := loadTodoList(filename); err != nil {
		fmt.Printf("Error loading to-do list: %v\n", err)
		return
	}

	task := a.ReadInput("Enter task: ")

	todoList = append(todoList, TodoItem{
		Task:     task,
		Complete: false,
	})

	if err := saveTodoList(filename); err != nil {
		fmt.Printf("Error saving to-do list: %v\n", err)
	} else {
		fmt.Println("Task added successfully!")
	}
}

// Mark a to-do item as complete

func MarkTodoComplete() {
	const filename = "todo_list.json"

	// Load to-do list from the file
	if err := loadTodoList(filename); err != nil {
		fmt.Printf("Error loading to-do list: %v\n", err)
		return
	}

	for i, item := range todoList {
		fmt.Printf("%d: %s\n", i+1, item.Task)
	}

	indexStr := a.ReadInput("Enter the number of the task to mark as complete: ")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(todoList) {
		fmt.Println("Invalid task number.")
		return
	}

	todoList[index-1].Complete = true

	if err := saveTodoList(filename); err != nil {
		fmt.Printf("Error saving to-do list: %v\n", err)
	} else {
		fmt.Println("Task marked as complete!")
	}
}
