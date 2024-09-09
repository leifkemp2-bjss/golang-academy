package main

import (
	"fmt"
	"sync"
	// "time"

	"academy.com/todoapp/todo"
)

func ex16(todoList ...todo.Todo){
	fmt.Println("Starting exercise 16")

	var wg sync.WaitGroup
	wg.Add(2)

	// var mu sync.Mutex

	ping := make(chan bool)

	go func(){
		defer wg.Done()

		for _, t := range todoList{
			// mu.Lock()
			fmt.Printf("%s\n", t.Contents)
			// mu.Unlock()
			ping <- true // Send the first channel ping, which activates the 2nd goroutine
			<- ping // Block here until we receive a channel ping from the 2nd goroutine
		}
	}()

	go func(){
		defer wg.Done()
		for _, t := range todoList{
			<- ping // Wait at the start until we receive a channel ping from the 1st goroutine
			// This receiver at the start of the loop ensures that the Contents goroutine is first
			// mu.Lock()
			fmt.Printf("%s\n", t.Status)
			// mu.Unlock()
			ping <- true // Send the second channel ping, which will allow the 1st goroutine to continue
		}
	}()

	wg.Wait()
	fmt.Println("Goroutines have finished")
}