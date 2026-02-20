package main

import "fmt"

type ToDo struct {
	IsComplete bool
	Name       string
	//ChangeCompletionStatus(bool)
}

func (t *ToDo) ChangeCompletion(completeStatus bool) {
	t.IsComplete = completeStatus
}

func main() {
	todo1 := ToDo{IsComplete: false, Name: "Learn Go!"}

	fmt.Println("TODO:", todo1.Name, "Completed?", todo1.IsComplete)

	todo1.ChangeCompletion(true)

	fmt.Println("TODO:", todo1.Name, "Completed?", todo1.IsComplete)
}
