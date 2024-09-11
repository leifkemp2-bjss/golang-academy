package todo

import (
	"reflect"
	"testing"
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

	_, err = testTodoList.CreateInMemory("", "")
	if err == nil {
		t.Error("empty contents are not accepted when creating a todo item")
	}

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

	_, err = testTodoList.CreateInMemory("Fifth todo item", "")
	if err == nil {
		t.Errorf("the list should be at capacity")
	}
	if err.Error() != "could not generate an unused ID for this Todo item" {
		t.Errorf("encountered an unexpected error: %s", err.Error())
	}
}

func TestUpdateInMemory(t *testing.T){

}

func TestDeleteInMemory(t *testing.T){

}