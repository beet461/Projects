package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func webpack(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("ls").Output()
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" : "hello"}`))
	fmt.Println(out, err)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/webpack", webpack)
	log.Fatal(http.ListenAndServe(":8080", r))
}
