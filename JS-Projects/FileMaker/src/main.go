package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func makego(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		exec.Command("ls").Output()
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message" : "hello"}`))
	}
	//	fmt.Println(out, err)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/makego", makego)
	log.Fatal(http.ListenAndServe(":8080", r))
}
