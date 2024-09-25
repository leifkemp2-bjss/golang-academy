package main

// This is version 1 of my web app, it has In Memory and SQL integration via Postgres, but is not
// thread safe.

import (
	// "bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"flag"

	"academy.com/todoapp/part2/flash"
	"academy.com/todoapp/todo"
	"academy.com/todoapp/part2/web/database"
)

var dir = "../../../files/todolist_web.json"
var todos *todo.TodoList
var useDatabase bool

func errorCheck(err error){
	if err != nil {
		log.Fatal(err) // log.fatal exits the program if there is an error
	}
}

func main(){
	flag.BoolFunc("db", "Uses the PostgreSQL database instead of the json files", func(s string) error {
		useDatabase = true
		return nil
	})

	flag.Parse()

	if useDatabase {
		database.DB = database.Connect()
		database.CreateTodosTable(database.DB)
	} else {
		todos, _ = InitialiseTodos()
	}
	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets"))))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/list", listTodosHandler)
	http.HandleFunc("/search/results", searchTodosHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/read/{id}", readTodosHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/update/{id}", updateHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/edit/{id}", editTodoHandler)
	http.HandleFunc("/delete/{id}", deleteHandler)
	fmt.Println("Running web app...")
	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
}

func rootHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"../html/main.html",
		"../html/header.html",
	))
	err := tmpl.Execute(writer, nil)
	errorCheck(err)
}

func listTodosHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"../html/list.html",
		"../html/header.html",
	))
	
	var err error
	if useDatabase{
		list, _ := database.ListTodos(database.DB)
		err = tmpl.Execute(writer, struct{
			Todos *todo.TodoList
			Flash string
		}{&list, ""})
	} else {
		err = tmpl.Execute(writer, struct{
			Todos *todo.TodoList
			Flash string
		}{todos, ""})
	}
	errorCheck(err)
}

func searchHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"../html/search.html",
		"../html/header.html",
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
	
	tmpl := template.Must(template.ParseFiles(
		"../html/searchresults.html",
		"../html/header.html",
	))

	var todosSearched []todo.Todo
	var err error
	if useDatabase{
		todosSearched, err = database.SearchForTodos(database.DB, contents, status)
	} else {
		todosSearched, err = todos.SearchInMemory(contents, status)
	}

	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, "/search", http.StatusFound)
		return
	}

	err = tmpl.Execute(writer, struct{
		Field string
		Todos []todo.Todo
	}{request.PathValue("contents"), todosSearched})
	errorCheck(err)
}

func readTodosHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	var t todo.Todo
	if useDatabase{
		t, _ = database.ReadTodo(database.DB, id)
	} else {
		t, err = todos.ReadInMemory(id)
		errorCheck(err)
	}

	tmpl := template.Must(template.ParseFiles(
		"../html/read.html",
		"../html/header.html",
	))
	err = tmpl.Execute(writer, t)
	errorCheck(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"../html/new.html",
		"../html/header.html",
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
	if useDatabase{
		database.InsertTodo(database.DB, contents, status)
	} else{
		_, err = todos.CreateInMemory(contents, status)
	}
	if err != nil{
		flash.SetFlash(writer, "message", []byte(err.Error()))
		http.Redirect(writer, request, "/new", http.StatusFound)
		return
	}

	if !useDatabase{
		Save(*todos)
	}

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func updateHandler(writer http.ResponseWriter, request *http.Request){
	flash, err := flash.GetFlash(writer, request, "message")
	errorCheck(err)

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	var t todo.Todo
	if useDatabase{
		t, _ = database.ReadTodo(database.DB, id)
	} else {
		t, err = todos.ReadInMemory(id)
		errorCheck(err)
	}
	
	tmpl := template.Must(template.ParseFiles(
		"../html/update.html",
		"../html/header.html",
	))

	err = tmpl.Execute(writer, struct{
		Todo todo.Todo
		Flash string
	}{t, string(flash)})

	errorCheck(err)
}

func editTodoHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	if useDatabase{
		database.UpdateTodo(database.DB, id, contents, status)
	} else {
		_, err = todos.UpdateInMemory(id, contents, status)
		if err != nil{
			flash.SetFlash(writer, "message", []byte(err.Error()))
			http.Redirect(writer, request, fmt.Sprintf("/update/%d", id), http.StatusFound)
			return
		}

		Save(*todos)
	}
	http.Redirect(writer, request, "/list", http.StatusFound)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	if useDatabase{
		database.DeleteTodo(database.DB, id)
	} else {
		err = todos.DeleteInMemory(id)
		errorCheck(err)
		Save(*todos)
	}

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func InitialiseTodos()(*todo.TodoList, error){
	todos := todo.TodoList{
		List: make(map[int]todo.Todo),
		MaxSize: 100,
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