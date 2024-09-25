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
	"strings"

	"academy.com/todoapp/todo"
	"academy.com/todoapp/part2/flash"
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
	http.HandleFunc("/search/results", searchTodosHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/read/{id}", readTodoHandler)	
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/update/{id}", updateHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/edit/{id}", editTodoHandler)
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

func searchHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"../../html/search.html",
		"../../html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Flash string
	}{string(flash)})
	errorCheck(err)
}

func searchTodosHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	contents = strings.TrimSpace(contents)
	status := request.FormValue("status")
	status = strings.TrimSpace(status)

	// this is checked in the frontend because without either of these filled, the backend thinks it will be a list
	if contents == "" && status == "" {
		flash.SetFlash(writer, "message", []byte("contents and status have not been provided, please provide at least one"))
		http.Redirect(writer, request, "/search", http.StatusFound)
		return
	}
	
	tmpl := template.Must(template.ParseFiles(
		"../../html/searchresults.html",
		"../../html/header.html",
	))

	var err error
	var jsonStr = []byte(fmt.Sprintf(`{"contents":"%s","status":"%s"}`, contents, status))
	req, err := http.NewRequest("GET", "http://localhost:8081/", bytes.NewBuffer(jsonStr))
	errorCheck(err)

	resp, err := client.Do(req)
	errorCheck(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	errorCheck(err)

	body := []todo.Todo{}
	json.Unmarshal(respBody, &body)

	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, "/search", http.StatusFound)
		return
	}

	err = tmpl.Execute(writer, struct{
		Field string
		Todos []todo.Todo
	}{request.PathValue("contents"), body})
	errorCheck(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"../../html/new.html",
		"../../html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Flash string
	}{string(flash)})
	errorCheck(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	var err error
	var jsonStr = []byte(fmt.Sprintf(`{"contents":"%s","status":"%s"}`, contents, status))
	req, err := http.NewRequest("POST", "http://localhost:8081/", bytes.NewBuffer(jsonStr))
	errorCheck(err)

	resp, err := client.Do(req)
	errorCheck(err)
	defer resp.Body.Close()

	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, "/new", http.StatusFound)
		return
	}

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func updateHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

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
		"../../html/update.html",
		"../../html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Todo todo.Todo
		Flash string
	}{body, string(flash)})

	errorCheck(err)
}

func editTodoHandler(writer http.ResponseWriter, request *http.Request){
	var err error
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d","contents":"%s","status":"%s"}`, id, contents, status))
	req, err := http.NewRequest("PUT", "http://localhost:8081/", bytes.NewBuffer(jsonStr))
	errorCheck(err)

	resp, err := client.Do(req)
	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, fmt.Sprintf("/update/%d", id), http.StatusFound)
		return
	}
	defer resp.Body.Close()

	http.Redirect(writer, request, "/list", http.StatusFound)
}