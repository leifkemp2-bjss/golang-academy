package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"academy.com/todoapp/part2/web/database"
)

var requestBuffer chan<- apiRequest

type apiRequest struct {
	verb     	string
	Id 			json.Number 	`json:"id"`
	Contents 	string 			`json:"contents"`
	Status 		string 			`json:"status"`
	response 	chan<- []byte
	err 		chan<- string
}

func startApi(ctx context.Context) <-chan struct{} {
	requests := make(chan apiRequest, 10)
	requestBuffer = requests
	done := make(chan struct{})
	// middleware
	// authentication
	srv := &http.Server{Addr: "localhost:8081"}

	http.HandleFunc("/", handle)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	// creates an actor channel and provides it with the context and requests channel
	loopDone := actor(requests, ctx)

	go func() {
		defer close(done)

		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Shutdown: " + err.Error())
		}
		<-loopDone
	}()

	return done
}

func actor(requests chan apiRequest, ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	// Shutdown
	go func() {
		<-ctx.Done()
		close(requests)
	}()

	go func() {
		defer close(done)
		<-processLoop(requests)
	}()

	return done
}

func processLoop(requests <-chan apiRequest) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for req := range requests {
			switch req.verb {
			case http.MethodGet:
				var result any
				var err error
				// if no id exists, it's either a list or a search, if an id exists, we get a specific one
				if req.Id == "" {
					if req.Contents == "" && req.Status == "" {
						result, err = database.ListTodos(database.DB)
						errorCheckChannel(err, req)
					} else {
						result, err = database.SearchForTodos(database.DB, req.Contents, req.Status)
						errorCheckChannel(err, req)
					}
				} else {
					id, err := req.Id.Int64()
					errorCheckChannel(err, req)

					result, err = database.ReadTodo(database.DB, int(id))
					errorCheckChannel(err, req)
				}
				value, err := json.Marshal(result)
				errorCheckChannel(err, req)

				req.response <- value
			case http.MethodPost:
				result, err := database.InsertTodo(database.DB, req.Contents, req.Status)
				errorCheckChannel(err, req)

				value, err := json.Marshal(result)
				errorCheckChannel(err, req)

				req.response <- value				
			case http.MethodDelete:
				id, err := req.Id.Int64()
				errorCheckChannel(err, req)

				err = database.DeleteTodo(database.DB, int(id))
				errorCheckChannel(err, req)

				req.response <- []byte{}
			case http.MethodPut:
				id, err := req.Id.Int64()
				errorCheckChannel(err, req)

				result, err := database.UpdateTodo(database.DB, int(id), req.Contents, req.Status)
				errorCheckChannel(err, req)

				value, err := json.Marshal(result)
				errorCheckChannel(err, req)
				req.response <- value
			}
			close(req.response)
			close(req.err)
		}
	}()
	return done
}

func errorCheckChannel(err error, req apiRequest){
	if err != nil {
		req.err <- err.Error()
	}
}

func handle(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}()

	responseChan := make(chan []byte, 1)
	errChan := make(chan string, 1)
	request := apiRequest{verb: r.Method, response: responseChan, err: errChan}

	body, err := io.ReadAll(io.Reader(r.Body))
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.Unmarshal(body, &request)

	select {
	case requestBuffer <- request:
	default:
		log.Println("request buffer full or shutdown")
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	select {
	case response, ok := <-responseChan:
		if r.Method != http.MethodGet && r.Method != http.MethodDelete {
			rw.WriteHeader(http.StatusOK)
			return
		}
	
		if ok {
			rw.WriteHeader(http.StatusOK)
			rw.Write(response)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	case err := <-errChan:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err))
	case <-time.After(5 * time.Second):
		rw.WriteHeader(http.StatusRequestTimeout)
		rw.Write([]byte("timed out after 5 seconds"))
	}
}