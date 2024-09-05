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

func (n name) String() string {
	if(n.middleName != ""){
		return fmt.Sprintf("%s %s %s", n.firstName, n.middleName, n.lastName)
	}
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

func assignment9(){
	reader := bufio.NewReader(os.Stdin)
	pl("Please enter a name with a first, middle (optional) and last name.")
	input, _ := reader.ReadString('\n')
	nameObj, err := createName(strings.TrimSpace(input))
	if err != nil {
		checkError(err)
	}
	pf("First Name: %s\n", nameObj.firstName)
	if nameObj.middleName != "" {
		 pf("Middle Name: %s\n", nameObj.middleName)
	}
	pf("Last Name: %s\n", nameObj.lastName)
}

func createName(input string)(*name, error){
	name := new(name)
	inputSplit := strings.Split(input, " ")

	if len(inputSplit) <= 1 {
		return nil, fmt.Errorf("this is not a valid name")
	}
	
	name.firstName = inputSplit[0]
	name.lastName = inputSplit[len(inputSplit)-1]

	if len(inputSplit) > 2{
		name.middleName = inputSplit[1]
	}

	return name, nil
}