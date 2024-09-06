package todo

import (
	"fmt"
	"encoding/json"
	"os"
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

// returns the list of Todo objects as a JSON object
func ListTodosAsJSON(todos ... Todo)([]byte, error){
	result, err := json.Marshal(todos)

	return result, err
}

func OutputTodosToJSONFile(dir string, todos ... Todo)(error){
	todosJSON, err := ListTodosAsJSON(todos...)

	if err != nil {
		return err
	}

	err = writeJSONToFile(dir, todosJSON)

	return err
}

func writeJSONToFile(dir string, json []byte)(error){
	f, err := os.Create(dir)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(json)

	return err
}