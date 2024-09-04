package main

import(
	"testing"
)

func TestValidNameFull(t *testing.T){
	nameTest := new(name)
	err := createName(nameTest, "John Peter Test")

	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if nameTest.firstName != "John" {
		t.Errorf("the first name has not been set properly, expecting John, got %s\n", nameTest.firstName)
	}
	if nameTest.middleName != "Peter" {
		t.Errorf("the middle name has not been set properly, expecting Peter, got %s\n", nameTest.middleName)
	}
	if nameTest.lastName != "Test" {
		t.Errorf("the last name has not been set properly, expecting Test, got %s\n", nameTest.lastName)
	}
}

func TestValidNameNoMiddle(t *testing.T){
	nameTest := new(name)
	err := createName(nameTest, "John Test")

	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if nameTest.firstName != "John" {
		t.Errorf("the first name has not been set properly, expecting John, got %s\n", nameTest.firstName)
	}
	if nameTest.lastName != "Test" {
		t.Errorf("the last name has not been set properly, expecting Test, got %s\n", nameTest.lastName)
	}

	if nameTest.middleName != "" {
		t.Errorf("the middle name has erroneously set, expecting an empty string, got %s\n", nameTest.middleName)
	}
}

func TestValidNameMultipleMiddle(t *testing.T){
	// the program should only create a name with one middle name
	nameTest := new(name)
	err := createName(nameTest, "John Peter Pete Test")

	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if nameTest.firstName != "John" {
		t.Errorf("the first name has not been set properly, expecting John, got %s\n", nameTest.firstName)
	}
	if nameTest.middleName != "Peter" {
		t.Errorf("the middle name has not been set properly, expecting Peter, got %s\n", nameTest.middleName)
	}
	if nameTest.lastName != "Test" {
		t.Errorf("the last name has not been set properly, expecting Test, got %s\n", nameTest.lastName)
	}
}

func TestInvalid(t *testing.T){
	nameTest := new(name)
	err := createName(nameTest, "John")

	if err == nil {
		t.Error("this name is invalid")
	}

	err = createName(nameTest, "")

	if err == nil {
		t.Error("this name is invalid")
	}
}