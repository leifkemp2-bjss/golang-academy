package main

import (
	// "bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"academy.com/todoapp/todo"
	"academy.com/todoapp/part2/flash"
)

var dir = "../../files/todolist_web.json"
var todos *todo.TodoList

func errorCheck(err error){
	if err != nil {
		log.Fatal(err) // log.fatal exits the program if there is an error
	}
}

func main(){
	todos, _ = InitialiseTodos()
	// errorCheck(err)
	// fmt.Println(todos)
	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/list", listTodosHandler)
	http.HandleFunc("/search/full/{contents}/{status}", searchTodosByFilterHandler)
	http.HandleFunc("/search/contents/{contents}", searchTodosByContentsHandler)
	http.HandleFunc("/search/status/{status}", searchTodosByStatusHandler)
	http.HandleFunc("/read/{id}", readTodosHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/update/{id}", updateHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/update_todo/{id}", updateTodoHandler)
	http.HandleFunc("/delete/{id}", deleteHandler)
	fmt.Println("Running web app...")
	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
}

func rootHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"html/main.html",
		"html/header.html",
	))
	err := tmpl.Execute(writer, nil)
	errorCheck(err)
}

func listTodosHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"html/list.html",
		"html/header.html",
	))

	err := tmpl.Execute(writer, struct{
		Todos *todo.TodoList
		Flash string
	}{todos, ""})
	errorCheck(err)
}

func searchTodosByFilterHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"html/searchresults.html",
		"html/header.html",
	))

	todosSearched := todos.SearchInMemoryByFilter(request.PathValue("contents"), request.PathValue("status"))

	err := tmpl.Execute(writer, struct{
		Field string
		Todos []todo.Todo
	}{request.PathValue("contents"), todosSearched})
	errorCheck(err)
}

func searchTodosByContentsHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"html/searchresults.html",
		"html/header.html",
	))

	todosSearched := todos.SearchInMemoryByContents(request.PathValue("contents"))

	err := tmpl.Execute(writer, struct{
		Field string
		Todos []todo.Todo
	}{request.PathValue("contents"), todosSearched})
	errorCheck(err)
}

func searchTodosByStatusHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"html/searchresults.html",
		"html/header.html",
	))

	todosSearched := todos.SearchInMemoryByStatus(request.PathValue("status"))

	err := tmpl.Execute(writer, struct{
		Field string
		Todos []todo.Todo
	}{request.PathValue("status"), todosSearched})
	errorCheck(err)
}

func readTodosHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)
	todo, err := todos.ReadInMemory(id)
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"html/read.html",
		"html/header.html",
	))
	err = tmpl.Execute(writer, todo)
	errorCheck(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"html/new.html",
		"html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Flash string
	}{string(flash)})
	errorCheck(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	_, err := todos.CreateInMemory(contents, status)
	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, "/new", http.StatusFound)
		return
	}

	Save(*todos)

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func updateHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)
	t, err := todos.ReadInMemory(id)
	errorCheck(err)
	
	tmpl := template.Must(template.ParseFiles(
		"html/update.html",
		"html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Todo todo.Todo
		Flash string
	}{t, string(flash)})

	errorCheck(err)
}

func updateTodoHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	_, err = todos.UpdateInMemory(id, contents, status)
	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, fmt.Sprintf("/update/%d", id), http.StatusFound)
		return
	}

	Save(*todos)

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	err = todos.DeleteInMemory(id)
	errorCheck(err)

	Save(*todos)

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func InitialiseTodos()(*todo.TodoList, error){
	todos := todo.TodoList{
		List: make(map[int]todo.Todo),
		MaxSize: 10,
	}

	err := todos.ReadTodosFromFileToMemory(dir)
	if err != nil {
		if _, ok := err.(*os.PathError); ok{
			fmt.Println("file does not exist, creating now")
			f, err := os.Create(dir)
			errorCheck(err)
			defer f.Close()
			f.Write([]byte(`[]`))
		} else {
			errorCheck(err)
		}
	}
	return &todos, nil
}

func Save(todos todo.TodoList){
	err := todos.SaveTodosFromMemoryToFile(dir)
	if err != nil {
		fmt.Println(err)
	}
}