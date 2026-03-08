package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type ToDo struct {
	IsComplete bool
	Name       string
}

// Marks a task complete or incomplete depending on completeStatus
func (t *ToDo) ChangeCompletion(completeStatus bool) {
	t.IsComplete = completeStatus
}

func main() {
	fmt.Printf("Welcome to Dylan's Golang To-Do List!\nType \"add\", \"list\", \"complete\", or \"remove\" to begin using the To-Do list.\nPress Ctrl+C to exit the program.\n")

	tasks := []ToDo{}
	// Used to get user input with spaces
	scanner := bufio.NewScanner(os.Stdin)

	// Main program loop
	for {
		fmt.Println("Commands: add, list, complete, remove")
		scanner.Scan()

		switch command := scanner.Text(); command {
		case "add":
			addCommand(scanner, &tasks)
		case "list":
			viewList(&tasks)
		case "remove":
			removeCommand(&tasks)
		case "complete":
			completeTask(tasks)
		}
		// fmt.Println("Tasks:", tasks)
	}
}

// Adds a task to the list
func addCommand(scanner *bufio.Scanner, list *[]ToDo) {
	fmt.Println("Create a task to add to your todo list...")
	scanner.Scan()
	// input := scanner.Text()
	// fmt.Println(input)
	*list = append(*list, ToDo{IsComplete: false, Name: scanner.Text()})
}

func viewList(list *[]ToDo) {
	// Display each task in the list
	for index, task := range *list {
		complete := "Complete"
		if !task.IsComplete {
			complete = "Not " + complete
		}

		fmt.Printf("%d. %s | %s\n", index, task.Name, complete)
	}
}

func removeCommand(list *[]ToDo) {
	fmt.Printf("Which task would you like to remove? Please provide an index (0-%d).\n", len(*list)-1)

	var input int
	fmt.Scanln(&input)

	newSlice := *list
	newSlice = slices.Delete(newSlice, input, input+1)
	*list = newSlice
}

func completeTask(list []ToDo) {
	fmt.Printf("Which task would you like to complete? Please provide an index (0-%d).\n", len(list)-1)

	var input int
	fmt.Scanln(&input)

	if input < len(list) {
		fmt.Println("Completing task", input)
		list[input].ChangeCompletion(true)
	}
}
