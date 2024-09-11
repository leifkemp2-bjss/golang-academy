package main

import (
	"fmt"
	"sync"
	// "time"
)

var value = 0
var mu sync.Mutex

func ex14() int {
	var wg sync.WaitGroup
	fmt.Println("Running exercise 14")

	// This exercise is designed to simulate a race condition
	wg.Add(2)

	go func(){
		for i := 0; i <= 10; i++{
			if i % 2 == 0 {
				value = i
				fmt.Printf("Setting value to %d\n", value)
			}
		}

		wg.Done()
	}()

	go func(){
		for i := 0; i <= 10; i++{
			if i % 2 != 0 {
				value = i
				fmt.Printf("Setting value to %d\n", value)
			}
		}

		wg.Done()
	}()

	wg.Wait()
	return value
}

func ex15_mutex() int{
	fmt.Println("Doing exercise 15 with mutexes")
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()

		for i := range 10{
			mu.Lock()
			value = 2 * i
			fmt.Printf("%d ", value)
			mu.Unlock()
		}
	}()

	go func(){
		defer wg.Done()
		for i := range 10{
			mu.Lock()
			value = 2 * i + 1
			fmt.Printf("%d ", value)
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println()
	fmt.Println("Goroutines have finished")
	fmt.Printf("Value is %d\n", value)
	
	return value
}

func ex15_channel(){
	fmt.Println("Doing exercise 15 with channels")
	var wg sync.WaitGroup
	evenChan := make(chan int)
	oddChan := make(chan int)
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i := range 10{
			evenChan <- 2 * i
		}
	}()

	go func(){
		defer wg.Done()
		for i := range 10{
			oddChan <- 2 * i + 1
		}
	}()

	for range 10{
		e := <-evenChan
		fmt.Printf("%d ", e)
		o := <-oddChan
		fmt.Printf("%d ", o)
	}

	wg.Wait()
	fmt.Println()
}