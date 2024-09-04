package main

import (
	"os"
	"strings"
	"unicode"
	"sort"
)

// sort functionality based on https://go.dev/play/p/twDrCxdYoi

// creates a custom type called ByCase, which takes an array of strings
type ByCase []string

// for a Sort function to work in Go, it needs the Len, Swap and Less functions, which are created here

func (s ByCase) Len() int { return len(s) }
func (s ByCase) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// this Less function accounts for the cases of the runes, uppercase runes are placed before lowercase runes
func (s ByCase) Less(i, j int) bool {
	iRunes := []rune(s[i])
	jRunes := []rune(s[j])

	// pick the shorter of the strings and grab the length from it
	max := len(iRunes)
	if max > len(jRunes){
		max = len(jRunes)
	}

	for x := 0; x < max; x++{
		ix := iRunes[x]
		jx := jRunes[x]

		lix := unicode.ToLower(ix)
		ljx := unicode.ToLower(jx)

		if lix != ljx {
			return lix < ljx
		}

		// the lowercase characters are the same, compare the originals

		if ix != jx {
			return ix < jx
		}
	}

	// we reach this point if the strings are identical
	return false
}
	
func assignment8(){
	arr := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}
	// Write the file first
	writeToFile(arr, "citynames")

	// Now read the file
	arr_from_file, err := readFile("citynames")
	checkError(err)
	pl(arr_from_file)

	// Arrange the list in alphabetical order

	sortByCase(arr_from_file)
	pl(arr_from_file)
}

func writeToFile(arr []string, filename string){
	f_w, err := os.Create("./files/" + filename)
	checkError(err)

	defer f_w.Close()

	for x := 0; x <= len(arr) - 1; x++{
		stringToWrite := arr[x]
		if x != len(arr) - 1 {
			stringToWrite += "\n"
		}

		size, err := f_w.WriteString(stringToWrite)
		checkError(err)
		pf("Wrote %d bytes to file\n", size)

		f_w.Sync()
	}
}

func readFile(filename string)([]string, error){
	f_r, err := os.ReadFile("./files/" + filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(f_r), "\n")
	return result, nil
}

func sortByCase(arr []string){
	sort.Sort(ByCase(arr))
}