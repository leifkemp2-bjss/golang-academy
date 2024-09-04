package main

import (
	"errors"
	"os"
	"testing"
	"reflect"
	"strings"
)

func init(){
	if _, err := os.Stat("./files/testfile"); err == nil {
		os.Remove("./files/testfile")
	}

	if _, err := os.Stat("./files/testfile"); err == nil {
		// The file still exists for some reason
		panic("testfile still exists, it should be deleted")
	}

	// write a testfile to read
	f_w, err := os.Create("./files/testfile2")
	arr := []string{"Test", "Test 2"}
	checkError(err)

	defer f_w.Close()

	for x := 0; x < len(arr); x++{
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

func shutdown(){
	if _, err := os.Stat("./files/testfile"); err == nil {
		os.Remove("./files/testfile")
	}
	if _, err := os.Stat("./files/testfile2"); err == nil {
		os.Remove("./files/testfile2")
	}
}

func TestMain(m *testing.M){
	m.Run() // force the tests to run sequentially
	shutdown()
}

func TestWriteFile(t *testing.T){
	arr := []string{"Test", "Test 2"}
	writeToFile(arr, "testfile")

	if _, err := os.Stat("./files/testfile"); err == nil {
		// The file has been created
	} else if errors.Is(err, os.ErrNotExist){
		t.Error("the file does not exist")
	} else {
		t.Error("the program hit an unexpected error trying to find the file")
	}

	result, err := os.ReadFile("./files/testfile")
	if err != nil {
		t.Error("the program hit an unexpected error trying to read the file")
	}

	expected := []string{"Test", "Test 2"}

	if !reflect.DeepEqual(strings.Split(string(result), "\n"), expected) {
		t.Errorf("the array contents have not been read properly, expecting %s, got %s", expected, result)
	}
}

func TestReadFile(t *testing.T){
	result, err := readFile("testfile2")

	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			t.Error("the file does not exist")
		} else {
			t.Error("the program hit an unexpected error trying to read the file")
		}
	}

	expected := []string{"Test", "Test 2"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("the array contents have not been read properly, expecting %s, got %s", expected, result)
	}
}