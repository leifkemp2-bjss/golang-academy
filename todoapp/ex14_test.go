package main

import(
	"testing"
)

func TestEx14(t *testing.T){
	got := ex14()
	want := 99
	if got != want {
		t.Errorf("program has not updated values properly, expected %d, got %d", want, got)
	}
}