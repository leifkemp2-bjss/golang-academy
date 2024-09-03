package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func assignment7(){
	for x := 0; x < 50; x++{
		dice1 := randRange(1, 6)
		dice2 := randRange(1, 6)

		sum := dice1 + dice2

		pl("Rolls: " + strconv.Itoa(dice1) + " " + strconv.Itoa(dice2) + " Sum: " + strconv.Itoa(sum))
		result, err := outcome(sum)
		if err != nil{
			panic(err)
		}
		pl(result)
	}
}

func randRange(min, max int) int {
	return rand.Intn(max+1-min)+min
}

func outcome(input int)(string, error){
	if input < 2 || input > 12{
		return "", fmt.Errorf("invalid outcome")
	}
	switch input{
	case 7, 11:
		return "NATURAL", nil
	case 2:
		return "SNAKE-EYES-CRAPS", nil
	case 3, 12:
		return "LOSS-CRAPS", nil
	default:
		return "NEUTRAL", nil
	}
}