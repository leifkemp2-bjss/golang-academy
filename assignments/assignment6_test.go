package main

import (
	"fmt"
	"testing"

	"time"

	"github.com/bearbin/go-age"
)

type MockAger struct{}

// This Ager runs the age function against a fixed date of 1st Jan, 2024 to make these tests future-proof.
func (m *MockAger) Age(birthDate time.Time) int{
	return age.AgeAt(birthDate, time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC))
}

func TestIsDateValid(t *testing.T){
	cases := []struct {
		day int
		month int
		year int
		outcome bool
	}{
		{day: 31, month: 1, year: 2001, outcome: true},
		{day: 29, month: 2, year: 2004, outcome: true},
		{day: 29, month: 2, year: 2000, outcome: true},
		{day: 29, month: 2, year: 1900, outcome: false},
		{day: 33, month: 1, year: 2001, outcome: false},
		{day: 0, month: 1, year: 2001, outcome: false},
		{day: -5, month: 1, year: 2001, outcome: false},
		{day: 31, month: 0, year: 2001, outcome: false},
		{day: 31, month: -5, year: 2001, outcome: false},
		{day: 31, month: 13, year: 2001, outcome: false},
		{day: 31, month: 1, year: -5, outcome: false},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing validity of %d-%d-%d", test.day, test.month, test.year), func(t *testing.T){
			got := isDateValid(test.day, test.month, test.year)
			if got != test.outcome {
				t.Errorf("got %t, want %t", got, test.outcome)
			}
		})
	}
}

func TestCalculateAge(t *testing.T){
	mockAger := &MockAger{}

	result, err := calculateAge(1, 1, 2023, mockAger)

	if err != nil {
		t.Error("the program has encountered an unexpected error")
	}
	if result != 1 {
		t.Errorf("the age has not been calculated correctly, expecting 1, got %d\n", result)
	}

	result, err = calculateAge(1, 9, 2023, mockAger)

	if err != nil {
		t.Error("the program has encountered an unexpected error")
	}
	if result != 0 {
		t.Errorf("the age has not been calculated correctly, expecting 0, got %d\n", result)
	}
}

func TestCalculateAgeInvalidDate(t *testing.T){
	mockAger := &MockAger{}

	_, err := calculateAge(0, 1, 2001, mockAger)

	if err == nil {
		t.Error("the program should have outputted an error for this invalid date")
	}
}