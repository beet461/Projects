package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func webpack(http.ResponseWriter, *http.Request) {
	fmt.Println("hello")
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/webpack", webpack)
	log.Fatal(http.ListenAndServe(":8000"), r)
}
