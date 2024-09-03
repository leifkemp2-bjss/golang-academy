package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Creating an alias of fmt.Println
var pl = fmt.Println

func main() {
	m := map[string]interface{}{
		"1": assignment1,
		"2": assignment2,
		"3": assignment3,
		"4": assignment4,
		"5": assignment5,
		"6": assignment6,
	}

	pl("Select an assignment (1-6)")

	reader := bufio.NewReader(os.Stdin)
	function, _ := reader.ReadString('\n')
	function = strings.TrimSpace(function)

	m[function].(func())()
}
