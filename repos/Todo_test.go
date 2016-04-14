package repos_test

import (
	m "nhaoday.com/models"
	"nhaoday.com/repos"

	"strings"
	"testing"
	"time"
)

func TestTodoList(t *testing.T) {
	v_todos, err := repos.TodoList()
	if err != nil {
		panic(err)
	}

	if v_todos == nil {
		t.Error("expected NOT nil")
	}

	if len(v_todos) != 7 {
		t.Error("expected number of rows equal to 7")
	}
}

func TestFindTodoById(t *testing.T) {
	v_todo, err := repos.FindTodoById(1)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(v_todo.Name, "Golang") {
		t.Error("expected Todo.Name contain Golang")
	}
}

func TestAddTodo(t *testing.T) {
	v_due := time.Date(2016, 9, 1, 0, 0, 0, 0, time.UTC)
	v_todo := &m.Todo{Name: "Build front-end by Redux and React.", Due: v_due}
	v_todo, err := repos.AddTodo(v_todo)
	if err != nil {
		panic(err)
	}
	if v_todo.Id == 0 {
		t.Error("expected NOT nil")
	}
}
