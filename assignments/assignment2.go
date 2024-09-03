package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkNameIsValid(name string){
	if(strings.TrimSpace(name) == ""){
		panic("No name has been provided.")
	}
}

func assignment2() {
	pl("Enter your first name")
	reader := bufio.NewReader(os.Stdin)
	firstName, _ := reader.ReadString('\n')
	checkNameIsValid(firstName)

	pl("Enter your middle name")
	middleName, _ := reader.ReadString('\n')

	pl("Enter your last name")
	lastName, _ := reader.ReadString('\n')
	checkNameIsValid(lastName)

	pl(buildName(strings.TrimSpace(firstName), strings.TrimSpace(middleName), strings.TrimSpace(lastName)))
}

func buildName(firstName, middleName, lastName string) (string, error){
	if firstName == "" {
		return "", fmt.Errorf("first name has not been provided")
	}
	if lastName == "" {
		return "", fmt.Errorf("last name has not been provided")
	}

	result := firstName
	if(middleName != ""){
		result += (fmt.Sprintf(" %s", middleName))
	}
	result += (fmt.Sprintf(" %s", lastName))
	return result, nil
}
