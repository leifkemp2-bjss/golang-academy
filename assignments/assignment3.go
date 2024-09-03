package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func assignment3() {
	for {
		pl("Enter an integer number")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			pl("This is not a valid input.")
		} else {
			if num >= 1 && num <= 10 {
				pl("Number is between 1 and 10.")
				break
			} else {
				pl("Number is not between 1 and 10.")
			}
		}
	}
}
