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

	pf("Sum of single digit numbers: %d\n", strconv.Itoa(singleDigitSum))
	pf("Sum of double digit numbers: %d\n", strconv.Itoa(doubleDigitSum))
	pf("Sum of triple digit numbers: %d\n", strconv.Itoa(tripleDigitSum))
	pf("Total: ", strconv.Itoa(singleDigitSum+doubleDigitSum+tripleDigitSum))
}

func sumNumbers(digits int) (sum int) {
	sum = 0
	count := 0
	reader := bufio.NewReader(os.Stdin)

	for count < 3 {
		pf("Please enter a %d digit number.\n", digits)
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
					pf("This input is not a %d digit number.\n", digits)
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
