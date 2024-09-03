package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func assignment5() {
	singleDigitSum := sumNumbers(1)
	doubleDigitSum := sumNumbers(2)
	tripleDigitSum := sumNumbers(3)

	pl("Sum of single digit numbers: " + strconv.Itoa(singleDigitSum))
	pl("Sum of double digit numbers: " + strconv.Itoa(doubleDigitSum))
	pl("Sum of triple digit numbers: " + strconv.Itoa(tripleDigitSum))
	pl("Total: " + strconv.Itoa(singleDigitSum+doubleDigitSum+tripleDigitSum))
}

func sumNumbers(digits int) (sum int) {
	sum = 0
	count := 0
	reader := bufio.NewReader(os.Stdin)

	for count < 3 {
		pl("Please enter a " + strconv.Itoa(digits) + " digit number.")
		input, err := reader.ReadString('\n')

		if err != nil {
			pl("This is not a valid input.")
			continue
		} else {
			inputTrimmed := strings.TrimSpace(input)
			num, err2 := strconv.Atoi(inputTrimmed)
			if err2 != nil {
				pl("This input is not an integer.")
				continue
			} else {
				if len(inputTrimmed) != digits {
					pl("This input is not a " + strconv.Itoa(digits) + " digit number.")
					continue
				} else {
					sum += num
					count++
				}
			}
		}
	}

	return
}
