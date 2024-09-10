package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"academy.com/todoapp/todo"
)

func main(){
	todos := todo.TodoList{}

	commandsList := map[string]string{
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
		commands := strings.Split(strings.TrimSpace(input), " ")

		switch commands[0] {
		case "help":
			for k, v := range commandsList {
				fmt.Printf("%s: %s\n", k, v)
			}
		case "read":
			if len(commands) < 2 {
				fmt.Println("id field has not been provided")
				continue
			}

			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("id field is invalid")
				continue
			}

			todo, err := todos.ReadInMemory(id)
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
}