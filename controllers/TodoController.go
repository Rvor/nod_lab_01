package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	u "nhaoday.com/common"
	m "nhaoday.com/models"
	d "nhaoday.com/repos"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos, err := d.TodoList()
	if err != nil {
		u.DisplayAppError(w, err, "Error when load todo list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		u.DisplayAppError(w, err, "Error when load todo list", http.StatusInternalServerError)
		return
		//panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, _ := strconv.Atoi(vars["todoId"])
	todo, err := d.FindTodoById(todoId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		u.DisplayAppError(w, err, "Wrong path", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		u.DisplayAppError(w, nil, "Error when create todo item", http.StatusMethodNotAllowed)
		return
	}
	var todo m.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		u.DisplayAppError(w, err, "Error when decode todo item", http.StatusInternalServerError)
		return
	}
	err = d.AddTodo(&todo)
	if err != nil {
		u.DisplayAppError(w, err, "Error when create todo item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		u.DisplayAppError(w, nil, "Error when update todo item", http.StatusMethodNotAllowed)
		return
	}

	var vars = mux.Vars(r)
	var Id, err = strconv.Atoi(vars["todoId"])
	if err != nil {
		u.DisplayAppError(w, err, "Error when update todo item", http.StatusInternalServerError)
		return
	}
	var todo m.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		u.DisplayAppError(w, err, "Invalid todo data.", http.StatusInternalServerError)
		return
	}
	err = d.UpdateTodo(Id, todo)
	if err != nil {
		u.DisplayAppError(w, err, "Error when update todo item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
