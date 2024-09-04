package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bearbin/go-age"
)

type Ager interface {
	Age(birthDate time.Time) int
}

type DefaultAger struct{}

func (d *DefaultAger) Age(birthDate time.Time) int{
	return age.Age(birthDate)
}

func assignment6() {
	ager := &DefaultAger{}
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

		age, err := calculateAge(dayI, monthI, yearI, ager)
		if err != nil { 
			pl(err)
			continue
		}

		pf("Age: %d", age)
		break
	}
}

func isDateValid(day, month, year int)(bool){
	dateString := fmt.Sprintf("%d-%.3s-%02d", year, time.Month(month), day)
	const shortForm = "2006-Jan-02"
	_, err := time.Parse(shortForm, dateString)
	return err == nil
}

func calculateAge(day, month, year int, ager Ager)(result int, err error){
	if !isDateValid(day, month, year) {
		return -1, fmt.Errorf("this date is not valid: %d-%d-%d", day, month, year)
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	pf("Date Constructed: %s\n", time.Time.String(date))

	result = ager.Age(date)

	return result, nil
}
