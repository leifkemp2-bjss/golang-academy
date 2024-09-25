package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"academy.com/todoapp/todo"
)

func errorCheck(err error){
	if err != nil {
		log.Fatal(err) // log.fatal exits the program if there is an error
	}
}

var client *http.Client

func main(){
	client = &http.Client{}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../../assets"))))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/list", listTodosHandler)
	http.HandleFunc("/delete/{id}", deleteTodoHandler)
	// http.HandleFunc("/search/results", searchTodosHandler)
	// http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/read/{id}", readTodoHandler)	
	// http.HandleFunc("/new", newHandler)
	// http.HandleFunc("/update/{id}", updateHandler)
	fmt.Println("Running web app...")
	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
}

func rootHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"../../html/main.html",
		"../../html/header.html",
	))
	err := tmpl.Execute(writer, nil)
	errorCheck(err)
}

func listTodosHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"../../html/list.html",
		"../../html/header.html",
	))

	req, err := http.NewRequest("GET", "http://localhost:8081/", nil)
	errorCheck(err)

	resp, err := client.Do(req)
	errorCheck(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	errorCheck(err)

	body := todo.TodoList{}
	json.Unmarshal(respBody, &body)

	err = tmpl.Execute(writer, struct{
		Todos *todo.TodoList
		Flash string
	}{&body, ""})
	errorCheck(err)
}

func readTodoHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d"}`, id))
	req, err := http.NewRequest("GET", "http://localhost:8081/", bytes.NewBuffer(jsonStr))
	errorCheck(err)

	resp, err := client.Do(req)
	errorCheck(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	errorCheck(err)

	body := todo.Todo{}
	json.Unmarshal(respBody, &body)

	tmpl := template.Must(template.ParseFiles(
		"../../html/read.html",
		"../../html/header.html",
	))
	err = tmpl.Execute(writer, body)
	errorCheck(err)
}

func deleteTodoHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d"}`, id))
	req, err := http.NewRequest("DELETE", "http://localhost:8081/", bytes.NewBuffer(jsonStr))
	errorCheck(err)

	resp, err := client.Do(req)
	errorCheck(err)
	defer resp.Body.Close()

	http.Redirect(writer, request, "/list", http.StatusFound)
}