package main

import (
	"fmt"
	"sync"
	// "time"
)

var value = 0

func ex14() int {
	fmt.Println("Running exercise 14/15")

	// ORIGINAL CODE
	// This exercise is designed to simulate a race condition
	
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func(){
	// 	for i := 0; i <= 20; i++{
	// 		if i % 2 == 0 {
	// 			value = i
	// 			fmt.Printf("Setting value to %d\n", value)
	// 		}
	// 	}

	// 	wg.Done()
	// }()

	// go func(){
	// 	for i := 0; i <= 20; i++{
	// 		if i % 2 != 0 {
	// 			value = i
	// 			fmt.Printf("Setting value to %d\n", value)
	// 		}
	// 	}

	// 	wg.Done()
	// }()

	// REFACTOR
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// chan1 := make(chan bool)

	// go func(){
	// 	for i := 0; i <= 20; i++{
	// 		if i % 2 == 0 {
	// 			mu.Lock()
	// 			value = i
	// 			mu.Unlock()
	// 			fmt.Printf("Setting value to %d\n", value)
	// 		}
	// 	}

	// 	wg.Done()
	// }()

	// go func(){
	// 	for i := 0; i <= 20; i++{
	// 		if i % 2 != 0 {
	// 			mu.Lock()
	// 			value = i
	// 			mu.Unlock()
	// 			fmt.Printf("Setting value to %d\n", value)
	// 		}
	// 	}

	// 	wg.Done()
	// }()

	ping := make(chan bool)

	go func(){
		defer wg.Done()

		for i := range 20{
			mu.Lock()
			value = 2 * i
			fmt.Printf("%d ", value)
			mu.Unlock()
			ping <- true // Send the first channel ping, which activates the 2nd goroutine
			<- ping // Block here until we receive a channel ping from the 2nd goroutine
		}
	}()

	go func(){
		defer wg.Done()
		for i := range 20{
			<- ping // Wait at the start until we receive a channel ping from the 1st goroutine
			// This receiver at the start of the loop ensures that the Contents goroutine is first
			mu.Lock()
			value = 2 * i + 1
			fmt.Printf("%d ", value)
			mu.Unlock()
			ping <- true // Send the second channel ping, which will allow the 1st goroutine to continue
		}
	}()

	wg.Wait()
	fmt.Println()
	fmt.Println("Goroutines have finished")
	fmt.Printf("Value is %d\n", value)
	
	return value
}