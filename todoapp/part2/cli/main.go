package main

import (
	"flag"
	"fmt"
	"os"

	"academy.com/todoapp/todo"
)

var dir = "../../files/todolist_cli.json"

var command string
var id int
var contents string
var status string

// TODO: explore using spawning processes in order to allow the program to continue running after executing a command - https://gobyexample.com/spawning-processes

func main(){
	todos, err := InitialiseTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	flag.StringVar(&command, "command", "", "Choose a command: help, read, list, create, update, delete")

	flag.BoolFunc("l", "Lists all todos", func(s string) error {
		command = "list"
		return nil
	})
	flag.BoolFunc("r", "Reads a todo by ID", func(s string) error {
		command = "read"
		return nil
	})
	flag.BoolFunc("d", "Deletes a todo by ID", func(s string) error {
		command = "delete"
		return nil
	})
	flag.BoolFunc("a", "Creates a todo, contents are required, status is optional", func(s string) error {
		command = "create"
		return nil
	})
	flag.BoolFunc("u", "Updates a todo with given ID, provided either contents or status are given", func(s string) error {
		command = "update"
		return nil
	})

	flag.IntVar(&id, "id", -1, "The ID of the todo")
	flag.IntVar(&id, "i", -1, "The ID of the todo (shorthand)")

	flag.StringVar(&contents, "contents", "", "The contents of the todo")
	flag.StringVar(&contents, "c", "", "The contents of the todo (shorthand)")

	flag.StringVar(&status, "status", "", "The status of the todo")
	flag.StringVar(&status, "s", "", "The status of the todo (shorthand)")

	flag.Parse()

	switch command{
	case "help":
		flag.PrintDefaults()
	case "read":
		if !IdValid(id) {
			return
		}

		todo, err := todos.ReadInMemory(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(todo)
	case "list":
		fmt.Print(todos.ListInMemory())
	case "create":
		if contents == "" {
			fmt.Println("content field has not been provided")
			return
		}

		i, err := todos.CreateInMemory(contents, status)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("New Todo created! Id: %d\n", i.Id)
		Save(todos)
	case "update":
		if !IdValid(id) {
			return
		}
		

		err := todos.UpdateInMemory(id, contents, status)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Todo ID %d updated!", id)
		Save(todos)
	case "delete":
		if !IdValid(id) {
			return
		}

		err := todos.DeleteInMemory(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Todo ID %d deleted!", id)
		Save(todos)
	}
}

func InitialiseTodos()(todo.TodoList, error){
	todos := todo.TodoList{
		List: make(map[int]todo.Todo),
		MaxSize: 100,
	}

	err := todos.ReadTodosFromFileToMemory(dir)
	if err != nil {
		
		if _, ok := err.(*os.PathError); ok{
			fmt.Println("file does not exist, creating now")
			f, err := os.Create(dir)

			if err != nil {
				fmt.Println(err)
				return todo.TodoList{}, err
			}
			defer f.Close()
			f.Write([]byte(`[]`))
		} else {
			fmt.Println(err)
			return todo.TodoList{}, err
		}
	}
	return todos, nil
}

func Save(todos todo.TodoList){
	err := todos.SaveTodosFromMemoryToFile(dir)
	if err != nil {
		fmt.Println(err)
	}
}

func IdValid(id int) bool {
	if id < 0 {
		fmt.Println("id field has not been provided")
	}
	return id >= 0
}