package todo

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testFileDir = "../files/testfile"

func setup(){
	if _, err := os.Stat(testFileDir); err == nil {
		os.Remove(testFileDir)
	}

	if _, err := os.Stat(testFileDir); err == nil {
		// The file still exists for some reason
		panic(testFileDir)
	}
}

func shutdown(){
	if _, err := os.Stat(testFileDir); err == nil {
		os.Remove(testFileDir)
	}

	if _, err := os.Stat(testFileDir); err == nil {
		// The file still exists for some reason
		panic(testFileDir)
	}
}

func TestMain(m *testing.M){
	setup()
	defer shutdown()
	m.Run()
}

func TestListTodos(t *testing.T){
	cases := []struct{
		todos []Todo
		want string
	}{
		{
			todos: []Todo{},
			want: "",
		},
		{
			todos: []Todo{{Contents: "create a test with only 1 Todo", Status: "todo"}},
			want: "0: create a test with only 1 Todo, Status: todo",
		},
		{
			todos: []Todo{
				{Contents: "create a test with multiple Todos", Status: "todo"},
				{Contents: "create multiple Todos for the test", Status: "completed"},
			},
			want: "0: create a test with multiple Todos, Status: todo\n1: create multiple Todos for the test, Status: completed",
		},
	}

	for _, test := range cases{
		t.Run(fmt.Sprintf("testing listing of %v", test.todos), func(t *testing.T) {
			got := ListTodos(test.todos...)

			if got != test.want{
				t.Errorf("expected %s got %s", test.want, got)
			}
		})
	}
}

func TestListTodosAsJSON(t *testing.T){
	cases := []struct{
		todos []Todo
		want []byte
	}{
		{
			todos: []Todo{},
			want: []byte(``),
		},
		{
			todos: []Todo{{Contents: "create a test with only 1 Todo", Status: "todo"}},
			want: []byte(`{"Contents": "create a test with only 1 Todo","Status": "todo"}`),
		},
	}

	for _, test := range cases{
		t.Run(fmt.Sprintf("testing listing of %v", test.todos), func(t *testing.T) {
			got, err := ListTodosAsJSON(test.todos...)

			if err != nil {
				t.Error("the program hit an unexpected error")
			}
			if reflect.DeepEqual(got, test.want){
				t.Errorf("expected %s got %s", test.want, got)
			}
		})
	}
}

func TestOutputTodosToJSONFile(t *testing.T){
	todoList := []Todo{
		{Contents: "create a test that outputs a Todo list to file", Status: "todo"},
		{Contents: "check the file exists", Status: "completed"},
		{Contents: "check the file's contents", Status: "inprogress"},
	}

	expected := []byte(
		`[{"Contents":"create a test that outputs a Todo list to file","Status":"todo"},`+ 
		`{"Contents":"check the file exists","Status":"completed"},`+
		`{"Contents":"check the file's contents","Status":"inprogress"}]`,
	)

	err := OutputTodosToJSONFile("../files/testfile.json", todoList...)
	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if _, err := os.Stat("../files/testfile"); err == nil {
		t.Error("the testfile.json does not exist")
	}

	result, err := os.ReadFile("../files/testfile.json")
	if err != nil {
		t.Error("the program hit an unexpected error trying to read the file")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("the array contents have not been read properly, expecting \n%s, got \n%s", expected, result)
	}
}