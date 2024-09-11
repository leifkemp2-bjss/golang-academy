package todo

import (
	"reflect"
	"testing"
	"fmt"
)

func TestReadInMemory(t *testing.T) {
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 5,
	}

	_, err := testTodoList.ReadInMemory(3)
	if err == nil {
		t.Error("the todo list does not contain a Todo item of id 3")
	}

	want := Todo{Id: 1, Contents: "Second todo item", Status: InProgress}
	got, err := testTodoList.ReadInMemory(1)
	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if !reflect.DeepEqual(got, want){
		t.Errorf("expecting %v, got %v", want, got)
	}
}

func TestListInMemory(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 5,
	}

	got := testTodoList.ListInMemory()
	want := "0: First todo item - Completed\n1: Second todo item - In Progress\n2: Third todo item - To Do\n"

	if got != want {
		t.Errorf("expecting %s, got %s", want, got)
	}
}

func TestCreateInMemory(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
		},
		MaxSize: 4,
	}
	var err error

	got, err := testTodoList.CreateInMemory("Third todo item", InProgress)
	want := Todo{Id: 2, Contents: "Third todo item", Status: InProgress}
	if err != nil {
		t.Error("the program hit an unexpected error")
	}
	if !reflect.DeepEqual(got, want){
		t.Errorf("expected %v, got %v", want, got)
	}

	got, err = testTodoList.CreateInMemory("Fourth todo item", "")
	want = Todo{Id: 3, Contents: "Fourth todo item", Status: ToDo}
	if err != nil {
		t.Error("the program hit an unexpected error")
	}
	if !reflect.DeepEqual(got, want){
		t.Errorf("expected %v, got %v", want, got)
	}

	if len(testTodoList.List) != 4{
		t.Error("the todo list should contain 4 items")
	}
}

func TestCreateInMemoryInvalid(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
		},
		MaxSize: 2,
	}
	var err error

	_, err = testTodoList.CreateInMemory("", "")
	if err == nil {
		t.Error("empty contents are not accepted when creating a todo item")
	}
	if err.Error() != "contents cannot be empty" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	}

	_, err = testTodoList.CreateInMemory("Third todo item", "")
	if err == nil {
		t.Errorf("the list should be at capacity")
	}
	if err.Error() != "could not generate an unused ID for this Todo item" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	}
}

func TestUpdateInMemory(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 4,
	}

	wants := []Todo{
		{Id: 0, Contents: "First todo item (updated)", Status: Completed},
		{Id: 1, Contents: "Second todo item", Status: Completed},
		{Id: 2, Contents: "Third todo item (update the status)", Status: InProgress},
	}

	err := testTodoList.UpdateInMemory(0, "First todo item (updated)", "")
	if err != nil {
		t.Error("encountered an unexpected error")
	}

	err = testTodoList.UpdateInMemory(1, "", Completed)
	if err != nil {
		t.Error("encountered an unexpected error")
	}

	err = testTodoList.UpdateInMemory(2, "Third todo item (update the status)", InProgress)
	if err != nil {
		t.Error("encountered an unexpected error")
	}

	for i := range wants{
		if !reflect.DeepEqual(testTodoList.List[i], wants[i]) {
			t.Errorf("expected %v, got %v", wants[i], testTodoList.List[i])
		}
	}
}

func TestUpdateInMemoryInvalid(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 4,
	}

	err := testTodoList.UpdateInMemory(5, "Mystery ID item", "")
	if err == nil {
		t.Error("this id is not valid")
	}
	if err.Error() != "item with id 5 does not exist" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	} 

	err = testTodoList.UpdateInMemory(2, "", "Nonsense Status")
	if err == nil {
		t.Error("this status is not valid")
	}
	if err.Error() != fmt.Sprintf("status is not valid, must be one of the following: %s, %s, %s", ToDo, InProgress, Completed) {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	} 

	err = testTodoList.UpdateInMemory(0, "", "")
	if err == nil {
		t.Error("at least one parameter must be filled")
	}
	if err.Error() != "content and status fields have not been provided, please provide at least one" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	}
}

func TestDeleteInMemory(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 4,
	}

	err := testTodoList.DeleteInMemory(0)
	if err != nil {
		t.Error("the program encountered an unexpected error")
	}

	if len(testTodoList.List) != 2{
		t.Error("the todo list should contain 2 items")
	}
}

func TestDeleteInMemoryInvalid(t *testing.T){
	testTodoList := TodoList{
		List: map[int]Todo{
			0: {Id: 0, Contents: "First todo item", Status: Completed},
			1: {Id: 1, Contents: "Second todo item", Status: InProgress},
			2: {Id: 2, Contents: "Third todo item", Status: ToDo},
		},
		MaxSize: 4,
	}

	err := testTodoList.DeleteInMemory(5)
	if err == nil {
		t.Error("the id does not exist in the todo list")
	}
	if err.Error() != "item with id 5 does not exist" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	} 
}