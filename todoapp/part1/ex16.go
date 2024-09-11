package main

import (
	"fmt"
	"sync"

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

	for range todoList {
		c := <- contentsChan
		fmt.Printf("%s: ", c)
		s := <- statusChan
		fmt.Printf("%s\n", s)
	}

	wg.Wait()
	fmt.Println("Goroutines have finished")
}