package todo

import "fmt"

type Todo struct {
	Contents string
	Status string
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, Status: %s", t.Contents, t.Status)
}

func ListTodos(todos ...Todo) string{
	result := ""
	for i, todo := range todos {
		result += fmt.Sprintf("%v", todo)
		if i != len(todos) - 1 {
			result += "\n"
		}
	}
	return result
}