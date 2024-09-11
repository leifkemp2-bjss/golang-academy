package main

import(
	"fmt"
	"academy.com/todoapp/todo"
)

func main(){
	todoList := []todo.Todo{
		{
			Id: 0,
			Contents: "foo bar todo",
			Status: todo.ToDo,
		},
		{
			Id: 1,
			Contents: "second foo bar todo",
			Status: todo.Completed,
		},
		{
			Id: 2,
			Contents: "third foo bar todo",
			Status: todo.InProgress,
		},
		{
			Id: 3,
			Contents: "fourth foo bar todo",
			Status: todo.Completed,
		},
	}
	
	err := todo.OutputTodosToJSONFile("../files/todolist.json", todoList...)

	if err != nil {
		fmt.Print(err)
	}

	readResult, _ := todo.ReadTodosFromFile("../files/todolist.json")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(todo.ListTodos(readResult...))

	ex14()

	ex15_channel()

	ex15_mutex()

	ex16(todoList...)
}