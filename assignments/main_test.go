package main

import(
	"testing"
)

// overrides the default main method provided in testing
// adding a shutdown method to remove all of the files created
// during the testing process
func TestMain(m *testing.M){
	setup_8()
	m.Run()
	shutdown_8()
}