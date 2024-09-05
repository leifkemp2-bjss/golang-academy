package main

import (
	"fmt"
	"testing"
)

func TestValidNames(t *testing.T){
	cases := []struct{
		input string
		name name
	}{
		{input: "John Peter Test", name: name{firstName: "John", middleName: "Peter", lastName: "Test",}},
		{input: "John Peter Pete Test", name: name{firstName: "John", middleName: "Peter", lastName: "Test",}},
		{input: "John Test", name: name{firstName: "John", lastName: "Test",}},
	}

	for _, test := range cases{
		t.Run(fmt.Sprintf("testing validity of %s", test.input), func(t *testing.T) {
			got, err := createName(test.input)
			if err != nil {
				t.Error("the program hit an unexpected error")
			}

			if(got.firstName != test.name.firstName){
				t.Errorf("firstName: expected %s, got %s", test.name.firstName, got.firstName)
			}
			if(got.middleName != test.name.middleName){
				t.Errorf("middleName: expected %s, got %s", test.name.middleName, got.middleName)
			}
			if(got.lastName != test.name.lastName){
				t.Errorf("lastName: expected %s, got %s", test.name.lastName, got.lastName)
			}
		})
	}
}

func TestInvalid(t *testing.T){
	_, err := createName("John")

	if err == nil {
		t.Error("this name is invalid")
	}

	_, err = createName("")

	if err == nil {
		t.Error("this name is invalid")
	}
}