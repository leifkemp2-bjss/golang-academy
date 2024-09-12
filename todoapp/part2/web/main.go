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
	fmt.Println(todos)
	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/list", listTodosHandler)
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

// func write(writer http.ResponseWriter, message string){
// 	_, err := writer.Write([]byte(message))	
// 	errorCheck(err)
// }

func rootHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"main.html",
		"header.html",
	))
	err := tmpl.Execute(writer, nil)
	errorCheck(err)
}

func listTodosHandler(writer http.ResponseWriter, request *http.Request){
	// template, err := template.ParseFiles("header.html", "main.html")
	tmpl := template.Must(template.ParseFiles(
		"list.html",
		"header.html",
	))
	// errorCheck(err)
	err := tmpl.Execute(writer, todos)
	errorCheck(err)
}

func readTodosHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Printf("Trying to read todo of id %s\n", request.PathValue("id"))
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)
	todo, err := todos.ReadInMemory(id)
	errorCheck(err)

	tmpl := template.Must(template.ParseFiles(
		"read.html",
		"header.html",
	))
	// errorCheck(err)
	err = tmpl.Execute(writer, todo)
	errorCheck(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request){
	tmpl := template.Must(template.ParseFiles(
		"new.html",
		"header.html",
	))

	err := tmpl.Execute(writer, nil)
	errorCheck(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	_, err := todos.CreateInMemory(contents, status)
	// errorCheck(err)
	if err != nil{
		
	}

	Save(*todos)

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func updateHandler(writer http.ResponseWriter, request *http.Request){
	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)
	todo, err := todos.ReadInMemory(id)
	errorCheck(err)
	
	tmpl := template.Must(template.ParseFiles(
		"update.html",
		"header.html",
	))

	err = tmpl.Execute(writer, todo)
	errorCheck(err)
}

func updateTodoHandler(writer http.ResponseWriter, request *http.Request){
	contents := request.FormValue("contents")
	status := request.FormValue("status")

	id, err := strconv.Atoi(request.PathValue("id"))
	errorCheck(err)

	err = todos.UpdateInMemory(id, contents, status)
	errorCheck(err)

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