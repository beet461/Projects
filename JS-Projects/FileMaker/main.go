package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func webpack(w http.ResponseWriter, *http.Request) {
	fmt.Println("hello")
	w
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/webpack", webpack)
	log.Fatal(http.ListenAndServe(":8080", r))
}
