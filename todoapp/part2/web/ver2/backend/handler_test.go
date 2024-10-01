package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
	"bytes"

	"academy.com/todoapp/part2/web/database"
	"academy.com/todoapp/todo"
)

var client *http.Client

func TestMain(m *testing.M){
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
	ctx, ctxDone := context.WithCancel(context.Background())
	done := createTestServer(ctx)
	
	m.Run()

	ctxDone()
	deleteTestDatabase()
	<- done
}

func TestList(t *testing.T){
	// reseed the database to account for if the tests are out of order
	seedTestDatabase()

	want := []todo.Todo{
		{Id: 1, Contents: "Test Todo 1", Status: "To Do"},
		{Id: 2, Contents: "Test Todo 2", Status: "In Progress"},
		{Id: 3, Contents: "Test Todo 3", Status: "Completed"},
	}

	responseBody := performRequest(t, "GET", nil, 200)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	body := todo.TodoList{}
	json.Unmarshal(respBody, &body)

	if len(body.List) != 3 {
		t.Error("expecting response with 3 todos")
	}

	for i, got := range body.List{
		if !reflect.DeepEqual(want[i-1], got) {
			t.Errorf("expected %v, got %v", want[i-1], got)
		}
	}
}

func TestGet(t *testing.T){
	want := todo.Todo{
		Id: 2, Contents: "Test Todo 2", Status: "In Progress",
	}

	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d"}`, 2))
	responseBody := performRequest(t, "GET", jsonStr, 200)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	body := todo.Todo{}
	json.Unmarshal(respBody, &body)

	if !reflect.DeepEqual(body, want){
		t.Errorf("expecting %v, got %v", want, body)
	}
}

func TestGetInvalid(t *testing.T){
	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d"}`, 999))
	responseBody := performRequest(t, "GET", jsonStr, 400)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	if string(respBody) != "sql: no rows in result set" {
		t.Errorf("got %s, want 'sql: no rows in result set'", string(respBody))
	}
}

func TestPost(t *testing.T){
	seedTestDatabase()

	want := todo.Todo{
		Id: 4,
		Contents: "This is my new todo",
		Status: "To Do",
	}

	var jsonStr = []byte(`{"status":"To Do","contents":"This is my new todo"}`)
	responseBody := performRequest(t, "POST", jsonStr, 200)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	var body int
	json.Unmarshal(respBody, &body)

	if body != 4 {
		t.Errorf("the update function should return the todo's ID after completing the update, got %d", body)
	}

	todoList, err := database.ListTodos(database.DB)
	if err != nil {
		t.Error(err)
	}
	if len(todoList.List) != 4 {
		t.Error("the todo list should have 4 todos after the post")
	}

	todoGot, err := database.ReadTodo(database.DB, 4)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(todoGot, want){
		t.Errorf("expecting %v, got %v", want, todoGot)
	}
}

func TestPostNoBody(t *testing.T){
	cases := []struct{
		jsonStr []byte
		want string
	}{
		{jsonStr: []byte(`{"status":"To Do"}`), want: "contents field has not been provided"},
		{jsonStr: []byte(`{"contents":"This is my new todo"}`), want: "status field has not been provided"},
	}

	for _, test := range cases {
		responseBody := performRequest(t, "POST", test.jsonStr, 400)
		defer responseBody.Close()

		respBody, err := io.ReadAll(responseBody)
		if err != nil {
			t.Error(err)
		}

		if string(respBody) != test.want {
			t.Errorf("got %s, want '%s'", string(respBody), test.want)
		}
	}
}

func TestPut(t *testing.T){
	seedTestDatabase()

	want := todo.Todo{
		Id: 1,
		Contents: "Update this todo",
		Status: "In Progress",
	}

	var jsonStr = []byte(`{"id":"1","status":"In Progress","contents":"Update this todo"}`)
	responseBody := performRequest(t, "PUT", jsonStr, 200)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	var body int
	json.Unmarshal(respBody, &body)

	if body != 1 {
		t.Errorf("the update function should return the todo's ID after completing the update, got %d", body)
	}

	todoGot, err := database.ReadTodo(database.DB, 1)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(todoGot, want){
		t.Errorf("expecting %v, got %v", want, todoGot)
	}
}

func TestPutInvalid(t *testing.T){
	var jsonStr = []byte(`{"id":"999","contents":"Non-existent todo (updated)","status":"In Progress"}`)
	responseBody := performRequest(t, "PUT", jsonStr, 400)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	if string(respBody) != "sql: no rows in result set" {
		t.Errorf("got %s, want 'sql: no rows in result set'", string(respBody))
	}
}

func TestPutNoBody(t *testing.T){
	var jsonStr = []byte(`{"id":"1"}`)
	responseBody := performRequest(t, "PUT", jsonStr, 400)
	defer responseBody.Close()

	respBody, err := io.ReadAll(responseBody)
	if err != nil {
		t.Error(err)
	}

	if string(respBody) != "content and status fields have not been provided" {
		t.Errorf("got %s, want 'content and status fields have not been provided'", string(respBody))
	}
}

func TestDelete(t *testing.T){
	seedTestDatabase()

	var jsonStr = []byte(`{"id":"1"}`)
	responseBody := performRequest(t, "DELETE", jsonStr, 200)
	responseBody.Close()

	todoList, err := database.ListTodos(database.DB)
	if err != nil {
		t.Error(err)
	}
	if len(todoList.List) != 2 {
		t.Error("the todo list should have 2 todos after the delete")
	}

	_, err = database.ReadTodo(database.DB, 1)
	if err == nil {
		t.Error("expecting an error when attempting to look for deleted todo of id 1")
	}
}

func TestDeleteInvalid(t *testing.T){
	seedTestDatabase()

	var jsonStr = []byte(`{"id":"999"}`)
	// deleting an invalid todo should do nothing, and return an OK status
	responseBody := performRequest(t, "DELETE", jsonStr, 200)
	responseBody.Close()

	todoList, err := database.ListTodos(database.DB)
	if err != nil {
		t.Error(err)
	}
	if len(todoList.List) != 3 {
		t.Error("the todo list should be unaffected by an invalid delete (have 3 todos)")
	}
}

func createTestServer(ctx context.Context) <- chan struct{}{
	database.DB = database.Connect()

	database.CreateTodosTable(database.DB, "todostest")
	seedTestDatabase()

	done := startApi(ctx, "localhost:8082")
	return done
}

func seedTestDatabase(){
	fmt.Println("Seeding DB")
	database.DB.Exec("TRUNCATE TABLE todostest RESTART IDENTITY")

	seedData := []struct{
		contents string
		status string
	}{
		{"Test Todo 1", "To Do"},
		{"Test Todo 2", "In Progress"},
		{"Test Todo 3", "Completed"},
	}

	for _, todo := range seedData{
		_, err := database.InsertTodo(database.DB, todo.contents, todo.status)
		if err != nil {
			log.Fatal(err)
		}
	}
	
}

func deleteTestDatabase(){
	database.DB.Exec("DROP TABLE todostest")
}

func performRequest(t *testing.T, verb string, jsonStr []byte, expectedStatusCode int) io.ReadCloser {
	req, err := http.NewRequest(verb, "http://localhost:8082/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("expecting a status code of %d, got %d", expectedStatusCode, resp.StatusCode)
	}
	return resp.Body
}