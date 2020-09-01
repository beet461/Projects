package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	key "github.com/beet461/Importable/KeyGen"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Data struct {
	Username string
	Password string
}

func makego(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		//get json input

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"username": "hello", "password": "helloo"}`))

		var data Data
		json.Unmarshal([]byte(body), &data)
		if data.Username == "nil" {
			fmt.Println("Username is nil")
		} else if data.Password == "" {
			fmt.Println("Password is nil")
		} else {
			fmt.Println("Everything is valid")
		}
		fmt.Println("Username =", data.Username)
		fmt.Println("Password =", data.Password)

		//db connect

		username := data.Username
		randkey := key.Gen()

		const (
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "shashi3969"
			dbname   = "FileMaker"
		)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Username =", username, "key = ", randkey)

		sqlStatement := fmt.Sprintf(`
	INSERT INTO api_keys (username, api_key)
	VALUES ('%v', '%v');`, data.Username, randkey)

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}
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
