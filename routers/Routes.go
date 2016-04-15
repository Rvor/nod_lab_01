package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	c "nhaoday.com/controllers"
)

type (
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	Routes []Route
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			HandlerFunc(r.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", c.PostIndex},
	Route{"TodoIndex", "GET", "/todos", c.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", c.TodoShow},
	Route{"AddTodo", "POST", "/todos", c.AddTodo},
	Route{"UpdateTodo", "PUT", "/todos/{todoId}", c.UpdateTodo},

	Route{"PostIndex", "GET", "/posts", c.PostIndex},
	Route{"PostShow", "GET", "/posts/{postId}", c.PostShow},
	Route{"AddPost", "POST", "/posts", c.AddPost},
	Route{"UpdatePost", "PUT", "/postos/{postId}", c.UpdatePost},
}
