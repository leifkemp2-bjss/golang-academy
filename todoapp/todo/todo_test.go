package todo

import (
	"fmt"
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
			want: "create a test with only 1 Todo, Status: todo",
		},
		{
			todos: []Todo{
				{Contents: "create a test with multiple Todos", Status: "todo"},
				{Contents: "create multiple Todos for the test", Status: "completed"},
			},
			want: "create a test with multiple Todos, Status: todo\ncreate multiple Todos for the test, Status: completed",
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