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

// Part 2 Functions

const (
	ToDo = "To Do"
	InProgress = "In Progress"
	Completed = "Completed"
)

type TodoList map[int]Todo

func (t TodoList) ReadInMemory (id int)(Todo, error){
	todo, err := t[id]
	if !err {
		return Todo{}, fmt.Errorf("Todo item with id %d not found", id)
	}
	return todo, nil
}

func (t TodoList) ListInMemory() string{
	result := ""
	for _, todo := range t {
		result += fmt.Sprintf("%v\n", todo)
	}
	return result
}

func (t TodoList) CreateInMemory(contents string)(Todo, error){
	if contents == "" {
		return Todo{}, fmt.Errorf("contents cannot be empty")
	}

	iters := 0
	for iters < 1000 {
		_, ok := t[iters]

		if !ok {
			// Found an Id that isn't in use
			break
		}
		iters++
	}

	if iters == 1000 {
		return Todo{}, fmt.Errorf("could not generate an unused ID for this Todo item")
	}

	todo := Todo{
		Id: iters,
		Contents: contents,
		Status: ToDo,
	}

	t[iters] = todo

	return todo, nil
}

func UpdateInMemory(){

}

func (t TodoList) DeleteInMemory(id int)(error){
	_, ok := t[id]
	if !ok {
		return fmt.Errorf("item with id %d does not exist", id)
	}
	delete(t, id)
	return nil
}