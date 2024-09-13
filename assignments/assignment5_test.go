package main

import (
	"testing"
	// "time"
	// "sync"
)

func TestProduceSum(t *testing.T){
	inputChan = make(chan int)
	outputChan = make(chan int)
	got := 0
	want := 10
	go produceSum(4)
	inputChan <- 1
	inputChan <- 2
	inputChan <- 3
	inputChan <- 4

	got = <- outputChan

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestProduceSumMultipleDigits(t *testing.T){
	inputChan = make(chan int)
	outputChan = make(chan int)
	got := 0
	want := 1111
	go produceSum(4)
	inputChan <- 1
	inputChan <- 10
	inputChan <- 100
	inputChan <- 1000

	got = <- outputChan

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}