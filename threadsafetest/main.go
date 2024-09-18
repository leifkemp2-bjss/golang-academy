package main

import (
	// "context"
	"fmt"
	"net/http"
	// "sync"
	// "log"

	// req "academy.com/threadsafetest/requestqueue"
)

// var rq req.RequestQueue
var requestChan chan apiRequest

// credit to https://github.com/labiraus/go-kvstore/blob/main/api.go for the apiRequest struct for thread safe server stuffs
type apiRequest struct{
	verb string
	response chan<- []byte // this is a write only channel
}

func main(){
	requestChan = make(chan apiRequest)
	fmt.Println("Started running :)")
	// rq = req.RequestQueue{
	// 	Chan: make(chan *http.Request),
	// }

	// http.HandleFunc("/", rootHandler)
	fmt.Println("Running web app...")
	go queueHandler()

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func queueHandler() string{
	fmt.Println("Queue handler beginning")
	for {
		fmt.Println("Listening for request chan")
		request := <- requestChan
		fmt.Println("Request chan called")
		request.response <- []byte(request.verb)
		close(request.response)
	}
}

func rootHandler(writer http.ResponseWriter, req *http.Request){
	fmt.Println("Handler called")

	responseChan := make(chan []byte)
	request := apiRequest{
		verb: req.Method,
		response: responseChan,
	}

	fmt.Println("Sending to request chan")
	requestChan <- request
	fmt.Println("Waiting for response chan")
	response, ok := <-responseChan

	if ok {
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}