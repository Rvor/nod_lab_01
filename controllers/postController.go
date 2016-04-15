package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	u "nhaoday.com/common"
	m "nhaoday.com/models"
	d "nhaoday.com/repos"
)

func PostIndex(w http.ResponseWriter, r *http.Request) {

	posts, err := d.PostList()

	if err != nil {
		u.DisplayAppError(w, err, "Error when load post list", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(PostsResource{Data: posts})

	if err != nil {
		u.DisplayAppError(w, err, "Error when load post list", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := strconv.Atoi(vars["postId"])
	post, err := d.FindPostById(postId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		u.DisplayAppError(w, err, "Data Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		u.DisplayAppError(w, nil, "Error when create post item", http.StatusMethodNotAllowed)
		return
	}
	var post m.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		u.DisplayAppError(w, err, "Error when decode post item", http.StatusInternalServerError)
		return
	}
	err = d.AddPost(&post)
	if err != nil {
		u.DisplayAppError(w, err, "Error when create post item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		u.DisplayAppError(w, nil, "Error when update post item", http.StatusMethodNotAllowed)
		return
	}

	var vars = mux.Vars(r)
	var Id, err = strconv.Atoi(vars["postId"])
	if err != nil {
		u.DisplayAppError(w, err, "Error when update post item", http.StatusInternalServerError)
		return
	}
	var post m.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		u.DisplayAppError(w, err, "Invalid post data.", http.StatusInternalServerError)
		return
	}
	err = d.UpdatePost(Id, post)
	if err != nil {
		u.DisplayAppError(w, err, "Error when update post item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
