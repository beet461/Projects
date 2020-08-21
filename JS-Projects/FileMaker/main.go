package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func webpack(http.ResponseWriter, *http.Request) {

}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1")
	api.HandlerFunc("/webpack", webpack)
}
