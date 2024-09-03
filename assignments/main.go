package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Creating an alias of fmt.Println
var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// Maps a string value to a function
	m := map[string]interface{}{
		"1": assignment1,
		"2": assignment2,
		"3": assignment3,
		"4": assignment4,
		"5": assignment5,
		"6": assignment6,
		"7": assignment7,
		"8": assignment8,
	}

	for {
		pl("Select an assignment (1-8)")

		reader := bufio.NewReader(os.Stdin)
		function, _ := reader.ReadString('\n')
		function = strings.TrimSpace(function)

		// Calls the function from the map
		value, exists := m[function]
		if(exists){
			value.(func())()
		}

		pl()
	}
}
