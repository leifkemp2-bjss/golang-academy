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
	todos := todo.TodoList{
		0: {Id: 0, Contents: "Read the list", Status: todo.InProgress},
		1: {Id: 1, Contents: "Add to the list", Status: todo.ToDo},
	}

	commandsList := map[string]string{
		"help": "Lists all available commands",
		"create": "Creates a new Todo item",
		"read {id}": "Reads a Todo item by id",
		"list": "Lists all Todo items",
		"update": "Updates a Todo item",
		"delete {id}": "Deletes a Todo item by id",
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Enter command (type help for list): ")
		input, _ := reader.ReadString('\n')
		inputTrimmed := strings.TrimSpace(input)
		commands := strings.Split(inputTrimmed, " ")

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
			fmt.Print(todos.ListInMemory())
		case "create":
			fmt.Print("Enter contents of Todo item: ")
			contents, _ := reader.ReadString('\n')
			
			i, err := todos.CreateInMemory(strings.TrimSpace(contents))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("New Todo created! Id: %d\n", i.Id)
		case "delete":
			if len(commands) < 2 {
				fmt.Println("id field has not been provided")
				continue
			}

			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("id field is invalid")
				continue
			}

			err = todos.DeleteInMemory(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		default:
			fmt.Println("Command not found. Use 'help' to list commands.")
		}
	}
}