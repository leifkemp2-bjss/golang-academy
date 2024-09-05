package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkNameIsValid(name string)(error){
	if(strings.TrimSpace(name) == ""){
		return fmt.Errorf("No name has been provided.")
	}
	return nil
}

func assignment2() {
	reader := bufio.NewReader(os.Stdin)
	firstName, middleName, lastName := "", "", ""
	for {
		pl("Enter your first name")
		firstName, _ = reader.ReadString('\n')
		if checkNameIsValid(firstName) != nil {
			continue
		}
		break
	}

	pl("Enter your middle name")
	middleName, _ = reader.ReadString('\n')

	for {
		pl("Enter your first name")
		lastName, _ = reader.ReadString('\n')
		if checkNameIsValid(lastName) != nil {
			continue
		}
		break
	}

	// if we get to this point the name is valid
	pl(buildName(strings.TrimSpace(firstName), strings.TrimSpace(middleName), strings.TrimSpace(lastName)))
}

func buildName(firstName, middleName, lastName string) (string){
	result := firstName
	if(middleName != ""){
		result += (fmt.Sprintf(" %s", middleName))
	}
	result += (fmt.Sprintf(" %s", lastName))
	return result
}
