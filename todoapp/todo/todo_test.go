package todo

import (
	"fmt"
	"reflect"
	"testing"
)

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