package main

import (
	"math/rand"
	"strconv"
)

func assignment7(){
	for x := 0; x < 50; x++{
		dice1 := randRange(1, 6)
		dice2 := randRange(1, 6)

		sum := dice1 + dice2

		pl("Rolls: " + strconv.Itoa(dice1) + " " + strconv.Itoa(dice2) + " Sum: " + strconv.Itoa(sum))
		pl(outcome(sum))
	}
}

func randRange(min, max int) int {
	return rand.Intn(max+1-min)+min
}

func outcome(input int)string{
	switch input{
	case 7, 11:
		return "NATURAL"
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 3, 12:
		return "LOSS-CRAPS"
	default:
		return "NEUTRAL"
	}
}