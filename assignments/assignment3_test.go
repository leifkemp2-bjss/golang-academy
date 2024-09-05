package main

import(
	"testing"
	"fmt"
)

func TestIsNumberInRange(t *testing.T){
	cases := []struct{
		num string
		outcome bool
	}{
		{num: "1", outcome: true},
		{num: "5", outcome: true},
		{num: "10", outcome: true},
		{num: "0", outcome: false},
		{num: "11", outcome: false},
		{num: "-1", outcome: false},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing validity of %s", test.num), func(t *testing.T){
			got, err := isNumberInRange(test.num)
			if err != nil {
				t.Error("the program has returned an unexpected error")
			}
			if got != test.outcome {
				t.Errorf("got %t, want %t", got, test.outcome)
			}
		})
	}
}

func TestInvalidInput(t *testing.T){
	_, err := isNumberInRange("Test")
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}
}