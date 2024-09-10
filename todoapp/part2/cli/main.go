package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"academy.com/todoapp/todo"
)

func main(){
	todos := todo.TodoList{}

	commands := map[string]string{
		"help": "Lists all available commands",
		"create": "Creates a new Todo item",
		"read {id}": "Reads a Todo item by id",
		"list": "Lists all Todo items",
		"update": "Updates a Todo item",
		"delete": "Deletes a Todo item",
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		command := strings.Split(strings.TrimSpace(input), " ")

		switch command[0] {
		case "help":
			for k, v := range commands {
				fmt.Printf("%s: %s\n", k, v)
			}
		case "read":
			todo, err := todos.ReadInMemory(0)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(todo)
		case "list":
			fmt.Print(todo.ListTodos())
		default:
			fmt.Println("Command not found. Use 'help' to list commands.")
		}

		fmt.Println()
	}

	fmt.Println(todos)
}