package main

import (
	"bufio"
	"os"
	"strings"
)

func assignment2() (result string) {
	pl("Enter your first name")
	reader := bufio.NewReader(os.Stdin)
	firstName, _ := reader.ReadString('\n')
	pl("Enter your middle name")
	middleName, _ := reader.ReadString('\n')
	pl("Enter your last name")
	lastName, _ := reader.ReadString('\n')
	result = "Hello " + strings.TrimSpace(firstName) + " " + strings.TrimSpace(middleName) + " " + strings.TrimSpace(lastName)
	pl(result)
	return
}
