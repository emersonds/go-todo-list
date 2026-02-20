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
	/*
		todo1 := ToDo{IsComplete: false, Name: "Learn Go!"}

		fmt.Println("TODO:", todo1.Name, "Completed?", todo1.IsComplete)

		todo1.ChangeCompletion(true)

		fmt.Println("TODO:", todo1.Name, "Completed?", todo1.IsComplete)
	*/
	tasks := make([]ToDo, 10)

	// Used to get user input with spaces
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Create a task to add to your todo list...")
	scanner.Scan()
	input := scanner.Text()
	tasks = append(tasks, ToDo{IsComplete: false, Name: input})

	fmt.Println("Task:", tasks[0].Name, "Completed:", tasks[0].IsComplete)
}
