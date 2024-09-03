package main

import (
	"bufio"
	"os"
	"strconv"
	"time"
	"strings"

	"github.com/bearbin/go-age"
)

func assignment6() {
	reader := bufio.NewReader(os.Stdin)

	for {
		pl("Input the day.")
		day, _ := reader.ReadString('\n')
		dayI, err := strconv.Atoi(strings.TrimSpace(day))
		if err != nil {
			pl("Invalid input for day.")
			continue
		}

		pl("Input the month.")
		month, _ := reader.ReadString('\n')
		monthI, err2 := strconv.Atoi(strings.TrimSpace(month))
		if err2 != nil {
			pl("Invalid input for month.")
			continue
		}

		pl("Input the year.")
		year, _ := reader.ReadString('\n')
		yearI, err3 := strconv.Atoi(strings.TrimSpace(year))
		if err3 != nil {
			pl("Invalid input for year.")
			continue
		}

		pl("Age: " + strconv.Itoa(calculateAge(dayI, monthI, yearI)))
		break
	}
}

func calculateAge(day, month, year int)(result int){
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	pf("Date Constructed: %s\n", time.Time.String(date))

	result = age.Age(date)

	return
}
