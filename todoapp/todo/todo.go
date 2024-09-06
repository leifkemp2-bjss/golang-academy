package todo

import (
	"fmt"
	"encoding/json"
)

type Todo struct {
	Contents string
	Status string
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, Status: %s", t.Contents, t.Status)
}

// returns the list of Todo objects in a printable and readable format
func ListTodos(todos ...Todo) string{
	result := ""
	for i, todo := range todos {
		result += fmt.Sprintf("%d: %v", i, todo)
		if i != len(todos) - 1 {
			result += "\n"
		}
	}
	return result
}

func ListTodosAsJSON(todos ... Todo)([]byte, error){
	result, err := json.Marshal(todos)

	return result, err
}