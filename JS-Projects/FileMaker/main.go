package main

import (
	"github.com/gorilla/mux"
)

func webpack() {

}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1")
	api.HandlerFunc("/webpack", webpack)
}
