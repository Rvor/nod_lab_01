package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	c "nhaoday.com/controllers"
)

func TestPostList(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/posts", c.PostIndex).Methods("GET")
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}