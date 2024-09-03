package main

import(
	"testing"
)

func TestValidOutcomes(t *testing.T){
	expectedOutputs := map[int]string{
		2:"SNAKE-EYES-CRAPS",
		3:"LOSS-CRAPS",
		4:"NEUTRAL",
		5:"NEUTRAL",
		6:"NEUTRAL",
		7:"NATURAL",
		8:"NEUTRAL",
		9:"NEUTRAL",
		10:"NEUTRAL",
		11:"NATURAL",
		12:"LOSS-CRAPS",
	}

	for k, v := range expectedOutputs{
		result, err := outcome(k)
		if err != nil{
			t.Error("the program has had an unexpected error")
		}
		if result != v{
			t.Errorf("incorrect outcome for num %d, expected %s but got %s", k, v, result)
		}
	} 
}

func TestInvalidOutcomes(t *testing.T){
	_, err := outcome(1)
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}

	_, err = outcome(13)
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}

	_, err = outcome(-1)
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}

	_, err = outcome(100)
	if err == nil{
		t.Error("the program has not returned an error when it should")
	}
}