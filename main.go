package main

import (
	"log"
	"net/http"

	r "nhaoday.com/routers"
)

func main() {
	router := r.NewRouter()

	log.Println("Listening at: 8080 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
