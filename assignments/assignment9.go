package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	firstName string
	middleName string
	lastName string
}

func assignment9(){
	reader := bufio.NewReader(os.Stdin)
	pl("Please enter a name with a first, middle (optional) and last name.")
	input, _ := reader.ReadString('\n')
	nameObj := new(name)
	err := createName(nameObj, strings.TrimSpace(input))
	if err != nil {
		checkError(err)
	}
	pf("First Name: %s\n", nameObj.firstName)
	if nameObj.middleName != "" {
		 pf("Middle Name: %s\n", nameObj.middleName)
	}
	pf("Last Name: %s\n", nameObj.lastName)
}

func createName(obj *name, input string)(error){
	inputSplit := strings.Split(input, " ")

	if len(inputSplit) <= 1 {
		return fmt.Errorf("this is not a valid name")
	}
	
	obj.firstName = inputSplit[0]
	obj.lastName = inputSplit[len(inputSplit)-1]

	if len(inputSplit) > 2{
		obj.middleName = inputSplit[1]
	}

	return nil
}