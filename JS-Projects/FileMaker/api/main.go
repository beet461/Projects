package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Username string
	Password string
}

func makego(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("body = ", body)
		strconv := string(body[:])
		fmt.Println("strconv = ", strconv)
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"username" : "hello", "password":"helloo" }`))
		//	json1 := `{"username" : "hello", "password":"helloo" }`
		var userpass Data
		json.Unmarshal([]byte(body), &userpass)
		fmt.Println(userpass.Username)
	}
	//	fmt.Println(out, err)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/makego", makego).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
