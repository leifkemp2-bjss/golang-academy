package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"academy.com/todoapp/part2/web/database"
)

var requestBuffer chan<- apiRequest

type apiRequest struct {
	verb     	string
	Id 			json.Number 	`json:"id"`
	Contents 	string 	`json:"contents"`
	Status 		string 	`json:"status"`
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
						if err != nil {
							req.err <- err.Error()
						}
					} else {
						result, err = database.SearchForTodos(database.DB, req.Contents, req.Status)
						if err != nil {
							req.err <- err.Error()
						}
					}
				} else {
					id, err := req.Id.Int64()
					if err != nil {
						req.err <- err.Error()
					}

					result, err = database.ReadTodo(database.DB, int(id))
					if err != nil {
						req.err <- err.Error()
					}
				}
				value, err := json.Marshal(result)
				if err != nil {
					req.err <- err.Error()
				}
				req.response <- value
			case http.MethodPost:
				result, err := database.InsertTodo(database.DB, req.Contents, req.Status)
				if err != nil {
					req.err <- err.Error()
				}
				value, err := json.Marshal(result)
				if err != nil {
					req.err <- err.Error()
				}
				req.response <- value				
			case http.MethodDelete:
				id, err := req.Id.Int64()
				if err != nil {
					req.err <- err.Error()
				}
				err = database.DeleteTodo(database.DB, int(id))
				if err != nil {
					req.err <- err.Error()
				}
				req.response <- []byte{}
			case http.MethodPut:
				id, err := req.Id.Int64()
				if err != nil {
					req.err <- err.Error()
				}
				result, err := database.UpdateTodo(database.DB, int(id), req.Contents, req.Status)
				if err != nil {
					req.err <- err.Error()
				}

				value, err := json.Marshal(result)
				if err != nil {
					req.err <- err.Error()
				}
				req.response <- value

			// case http.MethodPut:
			// case http.MethodPost:
			// case http.MethodPatch:
			// 	store[req.key] = req.value

			// case http.MethodGet:
			// 	value, ok := store[req.key]
			// 	if !ok {
			// 		log.Printf("could not find value for key %v", req.key)
			// 	} else {
			// 		req.response <- value
			// 	}
			}
			close(req.response)
		}
	}()
	return done
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
	}
}