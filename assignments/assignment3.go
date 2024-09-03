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

	isNumberInRange(input)
}

func isNumberInRange(input string)(bool, error){
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return false, fmt.Errorf("invalid input, must be an integer")
	} else {
		if num >= 1 && num <= 10 {
			pl("Number is between 1 and 10.")
			return true, nil
		} else {
			pl("Number is not between 1 and 10.")
			return false, nil
		}
	}
}
