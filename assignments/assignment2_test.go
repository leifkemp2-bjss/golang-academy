package main

import(
	"testing"
)

func TestFullNameCreated(t *testing.T){
	result, _ := buildName("John", "Peter", "Test")
	if result != "John Peter Test"{
		t.Error("full name has not been created")
	}
}

func TestMissingNames(t *testing.T){
	// A missing first name should report an error
	_, err := buildName("", "Peter", "Test")
	if err == nil {
		t.Error("this is not a valid name (missing first name)")
	}

	// A name should be valid without a middle name
	result, _ := buildName("John", "", "Test")
	if result != "John Test"{
		t.Error("this is a valid name (with no middle name)")
	}

	// A missing last name should report an error
	_, err = buildName("John", "Peter", "")
	if err == nil {
		t.Error("this is not a valid name (missing last name)")
	}
}