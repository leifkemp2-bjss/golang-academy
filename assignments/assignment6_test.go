package main

import (
	"testing"

	"github.com/bearbin/go-age"
	"time"
)

func init() {
	ageFunc = func(birthDate time.Time) int{
		return age.AgeAt(birthDate, time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC))
	}
}

func TestIsDateValid(t *testing.T){
	result := isDateValid(33, 1, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(0, 1, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(-5, 1, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(15, 0, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(15, -5, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(15, 13, 2001)

	if result {
		t.Error("this date is not valid")
	}

	result = isDateValid(15, 12, -5)

	if result {
		t.Error("this date is not valid")
	}
}

func TestCalculateAge(t *testing.T){
	result, err := calculateAge(1, 1, 2023)

	if err != nil {
		t.Error("the program has encountered an unexpected error")
	}
	if result != 1 {
		t.Errorf("the age has not been calculated correctly, expecting 1, got %d\n", result)
	}

	result, err = calculateAge(1, 9, 2023)

	if err != nil {
		t.Error("the program has encountered an unexpected error")
	}
	if result != 0 {
		t.Errorf("the age has not been calculated correctly, expecting 0, got %d\n", result)
	}
}

func TestCalculateAgeInvalidDate(t *testing.T){
	_, err := calculateAge(0, 1, 2001)

	if err == nil {
		t.Error("the program should have outputted an error for this invalid date")
	}
}