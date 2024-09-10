package todo

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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
		wants []string
	}{
		{
			todos: []Todo{},
			wants: []string{},
		},
		{
			todos: []Todo{{Id: 0, Contents: "create a test with only 1 Todo", Status: ToDo}},
			wants: []string{"create a test with only 1 Todo"},
		},
		{
			todos: []Todo{
				{Id: 0, Contents: "create a test with multiple Todos", Status: ToDo},
				{Id: 1, Contents: "create multiple Todos for the test", Status: Completed},
			},
			wants: []string{"create a test with multiple Todos", "create multiple Todos for the test"},
		},
	}

	for _, test := range cases{
		t.Run(fmt.Sprintf("testing listing of %v", test.todos), func(t *testing.T) {
			got := ListTodos(test.todos...)

			for _, want := range test.wants {
				if !strings.Contains(got, want) {
					t.Errorf("expecting output %s to contain %s", got, want)
				}
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
			todos: []Todo{{Id: 0, Contents: "create a test with only 1 Todo", Status: ToDo}},
			want: []byte(`{"Id":0,"Contents":"create a test with only 1 Todo","Status":"To Do"}`),
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
		{Id: 0, Contents: "create a test that outputs a Todo list to file", Status: ToDo},
		{Id: 1, Contents: "check the file exists", Status: Completed},
		{Id: 2, Contents: "check the file's contents", Status: InProgress},
	}

	expectedContents := []string{
		`"Id":0`,`"Contents":"create a test that outputs a Todo list to file"`,`"Status":"To Do"`,
		`"Id":1`,`"Contents":"check the file exists"`,`"Status":"Completed"`,
		`"Id":2`,`"Contents":"check the file's contents"`,`"Status":"In Progress"`,
	}

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

	for _, want := range expectedContents {
		if !strings.Contains(string(result), want) {
			t.Errorf("expecting output %s to contain %s", string(result), want)
		}
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