package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

var value = 0
var mu sync.Mutex

func ex14(n int) string {
	result := ""
	var wg sync.WaitGroup
	// fmt.Println("Running exercise 14")

	// This exercise is designed to simulate a race condition
	wg.Add(2)

	go func(){
		for i := range n{
			value = 2 * i
			result += fmt.Sprintf("%d", value)
			time.Sleep(1 * time.Millisecond)
		}

		wg.Done()
	}()

	go func(){
		for i := range n{
			value = 2 * i + 1
			result += fmt.Sprintf("%d", value)
			time.Sleep(1 * time.Millisecond)
		}

		wg.Done()
	}()

	wg.Wait()
	return result
}

func ex15_noconcurrency(n int) string{
	result := ""

	for i := range 2 * n {
		result += fmt.Sprintf("%d", i)
		time.Sleep(1 * time.Millisecond)
	}
	return result
}

func ex15_mutex(n int) string{
	// fmt.Println("Doing exercise 15 with mutexes")
	var wg sync.WaitGroup
	wg.Add(2)

	result := ""

	go func(){
		defer wg.Done()

		for i := range n{
			mu.Lock()
			value = 2 * i
			result += fmt.Sprintf("%d", value)
			// fmt.Printf("%d ", value)
			mu.Unlock()
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func(){
		defer wg.Done()
		for i := range n{
			mu.Lock()
			value = 2 * i + 1
			result += fmt.Sprintf("%d", value)
			mu.Unlock()
			time.Sleep(1 * time.Millisecond)
		}
	}()

	wg.Wait()	
	return result
}

func ex15_channel(n int)string{
	// fmt.Println("Doing exercise 15 with channels")
	var wg sync.WaitGroup
	wg.Add(2)

	numbersChan := make(chan int)

	go func(){
		defer wg.Done()
		for i := range n{
			numbersChan <- 2 * i
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func(){
		defer wg.Done()
		for i := range n{
			numbersChan <- 2 * i + 1
			time.Sleep(1 * time.Millisecond)
		}
	}()

	result := ""
	for range 2 * n{
		e := <-numbersChan
		result += fmt.Sprintf("%d ", e)
	}

	wg.Wait()
	close(numbersChan)
	return result
}

func ex15_manygoroutines(n int) string{
	numbersChan := make(chan int)

	for i := range 2 * n {
		go func(i int){
			numbersChan <- i
		}(i)
	}

	result := ""
	for range 2 * n{
		e := <-numbersChan
		result += fmt.Sprintf("%d ", e)
	}

	return result
}