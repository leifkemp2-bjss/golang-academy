package main

import (
	"fmt"
	"sync"
)

var value = 0

func ex14() int {
	// ORIGINAL CODE
	// This exercise is designed to simulate a race condition
	
	var wg sync.WaitGroup
	wg.Add(100)

	fmt.Println("Running exercise 14")

	for i := 0; i<50;i++ {
		go func (){
			defer wg.Done()
			value = i * 2
			fmt.Printf("Value is %d\n", value)
		}()
		go func (){
			defer wg.Done()
			value = i * 2 + 1
			fmt.Printf("Value is %d\n", value)
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines have finished")
	

	// REFACTOR

	return value
	
}