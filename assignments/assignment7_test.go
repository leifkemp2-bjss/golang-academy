package main

import(
	"testing"
	"fmt"
)

func TestValidOutcomes(t *testing.T){
	cases := []struct {
		num int
		outcome string
	}{
		{num: 2, outcome: "SNAKE-EYES-CRAPS"},
		{num: 3, outcome: "LOSS-CRAPS"},
		{num: 4, outcome: "NEUTRAL"},
		{num: 5, outcome: "NEUTRAL"},
		{num: 6, outcome: "NEUTRAL"},
		{num: 7, outcome: "NATURAL"},
		{num: 8, outcome: "NEUTRAL"},
		{num: 9, outcome: "NEUTRAL"},
		{num: 10, outcome: "NEUTRAL"},
		{num: 11, outcome: "NATURAL"},
		{num: 12, outcome: "LOSS-CRAPS"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing validity of %d", test.num), func(t *testing.T){
			got, err := outcome(test.num)
			if err != nil{
				t.Error("the program has had an unexpected error")
			}
			if got != test.outcome{
				t.Errorf("incorrect outcome for num %d, expected %s but got %s", test.num, test.outcome, got)
			}
		})
	}
}

func TestInvalidOutcomes(t *testing.T){
	cases := []int{1, 13, -1, 100}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing validity of %d", test), func(t *testing.T){
			_, err := outcome(test)
			if err == nil {
				t.Errorf("%d is not a valid input", test)
			}
		})
	}
}