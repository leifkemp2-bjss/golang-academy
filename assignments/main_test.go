package main

import(
	"testing"
	"time"
	"github.com/bearbin/go-age"
)

// This Ager runs the age function against a fixed date of 1st Jan, 2024 to make these tests future-proof.
func (m *MockAger) Age(birthDate time.Time) int{
	return age.AgeAt(birthDate, time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC))
}

// overrides the default main method provided in testing
// adding a shutdown method to remove all of the files created
// during the testing process
func TestMain(m *testing.M){
	setup_8()
	m.Run()
	shutdown_8()
}