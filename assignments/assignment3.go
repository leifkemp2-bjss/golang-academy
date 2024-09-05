package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func assignment3() {
	pl("Enter an integer number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result, err := isNumberInRange(strings.TrimSpace(input))
	checkError(err)
	
	if result {
		pl("Number is between 1 and 10.")
	} else {
		pl("Number is not between 1 and 10.")
	}
}

func isNumberInRange(input string)(bool, error){
	num, err := strconv.Atoi(input)
	if err != nil {
		return false, fmt.Errorf("invalid input, must be an integer")
	} else {
		return num >= 1 && num <= 10, nil
	}
}
