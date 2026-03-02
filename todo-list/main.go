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
	tasks := []ToDo{}
	// Used to get user input with spaces
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Create a task to add to your todo list...")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(input)
		tasks = append(tasks, ToDo{IsComplete: false, Name: input})

		fmt.Println("Tasks:", tasks)
	}
}
