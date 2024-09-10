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

	contentsChan := make(chan string)
	statusChan := make(chan string)

	go func(){
		defer wg.Done()
		for _, t := range todoList{
			contentsChan <- t.Contents
		}
	}()

	go func(){
		defer wg.Done()
		for _, t := range todoList{
			statusChan <- t.Status
		}
	}()

	for _, todo := range todoList {
		c := <- contentsChan
		fmt.Println(c)
		s := <- statusChan
		fmt.Println(s)
		fmt.Printf("Todo %d: %s - %s\n", todo.Id, c, s)
	}

	wg.Wait()
	fmt.Println("Goroutines have finished")
}