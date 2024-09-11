package todo

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
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

type TodoList struct {
	List map[int]Todo
	MaxSize int
}

func (t TodoList) ReadInMemory (id int)(Todo, error){
	todo, err := t.List[id]
	if !err {
		return Todo{}, fmt.Errorf("Todo item with id %d not found", id)
	}
	return todo, nil
}

func (t TodoList) ListInMemory() string{
	result := ""
	for _, todo := range t.List {
		result += fmt.Sprintf("%v\n", todo)
	}
	return result
}

func (t TodoList) CreateInMemory(contents string, status string)(Todo, error){
	if contents == "" {
		return Todo{}, fmt.Errorf("contents cannot be empty")
	}

	iters := 0
	for iters < t.MaxSize {
		_, ok := t.List[iters]

		if !ok {
			// Found an Id that isn't in use
			break
		}
		iters++
	}

	if iters == t.MaxSize {
		return Todo{}, fmt.Errorf("could not generate an unused ID for this Todo item")
	}

	todo := Todo{
		Id: iters,
		Contents: contents,
		Status: ToDo,
	}

	if status != "" {
		if status != ToDo && status != InProgress && status != Completed {
			return Todo{}, fmt.Errorf("status is not valid, must be one of the following: %s, %s, %s", ToDo, InProgress, Completed)
		}
		todo.Status = status
	}

	t.List[iters] = todo

	return todo, nil
}

func (t TodoList) UpdateInMemory(id int, contents string, status string)(error){
	_, ok := t.List[id]
	if !ok {
		return fmt.Errorf("item with id %d does not exist", id)
	}

	todo := Todo{
		Id: id,
		Contents: t.List[id].Contents,
		Status: t.List[id].Status,
	}

	if contents != "" {
		todo.Contents = contents
	}
	if status != "" {
		if status != ToDo && status != InProgress && status != Completed {
			return fmt.Errorf("status is not valid, must be one of the following: %s, %s, %s", ToDo, InProgress, Completed)
		}
		todo.Status = status
	}

	t.List[id] = todo
	return nil
}

func (t TodoList) DeleteInMemory(id int)(error){
	_, ok := t.List[id]
	if !ok {
		return fmt.Errorf("item with id %d does not exist", id)
	}
	delete(t.List, id)
	return nil
}

func (t TodoList)ReadTodosFromFileToMemory(dir string)(error){
	f_r, err := os.ReadFile(dir)
	if err != nil {
		return err
	}

	result := []Todo{}
	err = json.Unmarshal(f_r, &result)
	
	for _, r := range result {
		t.List[r.Id] = r
	}

	return err
}

func (t *TodoList)SaveTodosFromMemoryToFile(dir string)(error){
	values := slices.Collect(maps.Values(t.List))
	OutputTodosToJSONFile(dir, values...)
	return nil
}