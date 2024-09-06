package main

import(
	"fmt"
	"academy.com/todoapp/todo"
)

func main(){
	fmt.Println("foo")	

	fmt.Print(todo.ListTodos(todo.Todo{
			Contents: "foo bar todo",
			Status: "todo",
		},
		todo.Todo{
			Contents: "second foo bar todo",
			Status: "completed",
		},
	))
}