package main

import (
	"bufio"
	"fmt"
	"os"
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
	tasks := []ToDo{}
	// Used to get user input with spaces
	scanner := bufio.NewScanner(os.Stdin)

	// Main program loop
	for {
		fmt.Println("Commands: add, list, complete")
		scanner.Scan()

		switch command := scanner.Text(); command {
		case "add":
			addCommand(scanner, &tasks)
		}
		for index, task := range tasks {
			fmt.Println(index, task)
		}
		fmt.Println("Tasks:", tasks)
	}
}

// Adds a task to the list
func addCommand(scanner *bufio.Scanner, list *[]ToDo) {
	fmt.Println("Create a task to add to your todo list...")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
	*list = append(*list, ToDo{IsComplete: false, Name: input})
}
