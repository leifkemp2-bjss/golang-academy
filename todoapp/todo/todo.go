package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Id int
	Contents string
	Status string
}

func (t Todo) String() string {
	return fmt.Sprintf("%d: %s - %s", t.Id, t.Contents, t.Status)
}

// returns the list of Todo objects in a printable and readable format
func ListTodos(todos ...Todo) string{
	result := ""
	for _, todo := range todos {
		result += fmt.Sprintf("%v\n", todo)
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

func ReadTodosFromFile(dir string)([]Todo, error){
	f_r, err := os.ReadFile(dir)
	if err != nil {
		return nil, err
	}

	result := []Todo{}
	
	err = json.Unmarshal(f_r, &result)

	return result, err
}

