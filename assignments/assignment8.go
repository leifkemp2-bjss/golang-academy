package main

import (
	"os"
	"strings"
)
	
func assignment8(){
	arr := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}
	pl("foo")

	// Write the file first
	f_w, err := os.Create("./files/citynames")
	checkError(err)

	defer f_w.Close()

	for x := 0; x < len(arr); x++{
		size, err := f_w.WriteString(arr[x] + "\n")
		checkError(err)
		pf("Wrote %d bytes to file\n", size)

		f_w.Sync()
	}

	// Now read the file
	f_r, err := os.ReadFile("./files/citynames")
	checkError(err)
	arr_from_file := strings.Split(string(f_r), "\n")
	pl(arr_from_file)

	// Arrange the list in alphabetical order
}