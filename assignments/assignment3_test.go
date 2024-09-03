package main

import(
	"testing"
)

func TestValidNumber(t *testing.T){
	result, err := isNumberInRange("1")
	if err != nil{
		t.Error("the program has returned an error when it shouldn't")
	}
	if result != true{
		t.Error("the program has returned false, expecting true")
	}
}

func TestInvalidNumber(t *testing.T){
	result, err := isNumberInRange("-1")
	if err != nil{
		t.Error("the program has returned an error when it shouldn't")
	}
	if result != false{
		t.Error("the program has returned true, expecting false")
	}
}

func TestInvalidInp(t *testing.T){
	_, err := isNumberInRange("Test")
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}
}