package main

import(
	"fmt"
	"academy.com/todoapp/todo"
)

func main(){
	todoList := []todo.Todo{
		{
			Contents: "foo bar todo",
			Status: "todo",
		},
		{
			Contents: "second foo bar todo",
			Status: "completed",
		},
		{
			Contents: "third foo bar todo",
			Status: "inprogress",
		},
		{
			Contents: "fourth foo bar todo",
			Status: "completed",
		},
	}
	
	err := todo.OutputTodosToJSONFile("./files/todolist.json", todoList...)

	if err != nil {
		fmt.Print(err)
	}

	readResult, _ := todo.ReadTodosFromFile("./files/todolist.json")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(todo.ListTodos(readResult...))

	ex14()

	ex16(todoList...)
}