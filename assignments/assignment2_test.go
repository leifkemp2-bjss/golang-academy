package main

import(
	"testing"
)

func TestBuildName(t *testing.T){
	result := buildName("John", "Peter", "Test")
	if result != "John Peter Test"{
		t.Error("full name has not been created")
	}

	result = buildName("John", "", "Test")
	if result != "John Test"{
		t.Error("full name (without middle name) has not been created")
	}
}

func TestCheckNameIsValid(t *testing.T){
	if checkNameIsValid("John") != nil {
		t.Error("this is a valid name")
	}

	if checkNameIsValid("") == nil {
		t.Error("this is not a valid name")
	}
}