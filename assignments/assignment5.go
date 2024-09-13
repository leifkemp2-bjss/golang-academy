package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var inputChan chan int
var outputChan chan int
var sum int
func assignment5() {
	sum = 0
	inputChan = make(chan int)
	outputChan = make(chan int)
	go produceSum(9)
	listenForInputs(3, 1)
	listenForInputs(3, 2)
	listenForInputs(3, 3)
	
	sum=<-outputChan
	pl(sum)
}

func produceSum(calls int)int{
	sum := 0
	for range calls{
		val := <- inputChan
		sum += val
	}
	outputChan <- sum
	return sum
}

func listenForInputs(i int, digits int){
	reader := bufio.NewReader(os.Stdin)
	count := 0
	for count < i {
		pf("Please enter a %d digit number: \n", digits)
		input, err := reader.ReadString('\n')
		
		if err != nil {
			pl("This is not a valid input.")
			continue
		} else {
			inputTrimmed := strings.TrimSpace(input)
			num, err := strconv.Atoi(inputTrimmed)
			
			if err != nil {
				pl("This input is not an integer.")
				continue
			} else {
				if len(inputTrimmed) != digits {
					pf("This input is not a %d digit number.\n", digits)
					continue
				} else {
					inputChan <- num
					count++
				}
			}
		}
	} 
}