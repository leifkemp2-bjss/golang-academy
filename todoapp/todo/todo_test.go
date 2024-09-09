package todo

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testFileDirs = []string{"../files/testfile.json", "../files/testfile_read.json"}

func setup(){
	for _, testFileDir := range testFileDirs{
		if _, err := os.Stat(testFileDir); err == nil {
			os.Remove(testFileDir)
		}

		if _, err := os.Stat(testFileDir); err == nil {
			// The file still exists for some reason
			panic(fmt.Sprintf("the %s file still exists when it should be removed", testFileDir))
		}
	}

	// now that everything is torn down, recreate testfile_read.json and populate it
	f, _ := os.Create(testFileDirs[1])
	defer f.Close()
	f.Write([]byte(`[{"Contents":"test reading files","Status":"todo"}]`))
}

func shutdown(){
	for _, testFileDir := range testFileDirs{
		if _, err := os.Stat(testFileDir); err == nil {
			os.Remove(testFileDir)
		}

		if _, err := os.Stat(testFileDir); err == nil {
			// The file still exists for some reason
			panic(fmt.Sprintf("the %s file still exists when it should be removed", testFileDir))
		}
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

	err := OutputTodosToJSONFile(testFileDirs[0], todoList...)
	if err != nil {
		t.Error("the program hit an unexpected error")
	}

	if _, err := os.Stat(testFileDirs[0]); err != nil {
		t.Error("the testfile.json does not exist")
	}

	result, err := os.ReadFile(testFileDirs[0])
	if err != nil {
		t.Error("the program hit an unexpected error trying to read the file")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("the array contents have not been read properly, expecting \n%s, got \n%s", expected, result)
	}
}

func TestReadTodosFromFile(t *testing.T){
	if _, err := ReadTodosFromFile("../files/testfile_doesntexist.json"); err == nil {
		t.Error("the program did not return an error about the missing file")
	}

	if _, err := os.Stat(testFileDirs[1]); err != nil {
		t.Errorf("the %s file does not exist", testFileDirs[1])
	}

	expected := []Todo{
		{
			Contents: "test reading files",
			Status: "todo",
		},
	}

	result, err := ReadTodosFromFile(testFileDirs[1])
	if err != nil {
		t.Error("the program hit an unexpected error trying to read the file")
	}

	if !reflect.DeepEqual(expected, result){
		t.Errorf("expecting %v, got %v", expected, result)
	}
}