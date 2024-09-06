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
	}
	fmt.Println("foo")	

	fmt.Println(todo.ListTodos(todoList...))

	todoJSON, _ := todo.ListTodosAsJSON(todoList...)

	fmt.Println(string(todoJSON))

	err := todo.OutputTodosToJSONFile("./files/todolist.json", todoList...)

	if err != nil {
		fmt.Print(err)
	}
}