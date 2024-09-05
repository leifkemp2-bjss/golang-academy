package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkNameIsValid(name string)(error){
	if(name == ""){
		return fmt.Errorf("no name has been provided")
	}
	return nil
}

func assignment2() {
	reader := bufio.NewReader(os.Stdin)
	firstName, middleName, lastName := "", "", ""
	for {
		pl("Enter your first name")
		firstName, _ = reader.ReadString('\n')
		firstName = strings.TrimSpace(firstName)
		if checkNameIsValid(firstName) != nil {
			continue
		}
		break
	}

	pl("Enter your middle name")
	middleName, _ = reader.ReadString('\n')
	middleName = strings.TrimSpace(middleName)

	for {
		pl("Enter your first name")
		lastName, _ = reader.ReadString('\n')
		lastName = strings.TrimSpace(lastName)
		if checkNameIsValid(lastName) != nil {
			continue
		}
		break
	}

	// if we get to this point the name is valid
	pl(buildName(firstName, middleName, lastName))
}

func buildName(firstName, middleName, lastName string) (string){
	result := firstName
	if(middleName != ""){
		result += (fmt.Sprintf(" %s", middleName))
	}
	result += (fmt.Sprintf(" %s", lastName))
	return result
}
