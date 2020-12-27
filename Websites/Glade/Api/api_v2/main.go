package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Data ..
type Data struct {
	Email    string
	Username string
	Password string
}

type respondValues struct {
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func unmarshalData(w http.ResponseWriter, r *http.Request) Data {
	var data Data
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return data
	}
	if r.Method == "POST" {
		//Get json input
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		//Success and enabling cors
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":"true"}`))

		//Unmarshaling data and storing it in data
		json.Unmarshal([]byte(body), &data)
		return data
	}
	return data
}

func doQuery(sqlStatement, reqType string) (scanValues [3]string) {
	//Defining the values for connecting to the database
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "shashi3969"
		dbname   = "Glade"
	)

	//Database values
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Connecting/opening to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//Closing the database at the end
	defer db.Close()

	//Pinging the database for any errors
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var scanResult [3]string

	//Query the statement in the database
	queryResult, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Query error : ", err)
	}

	if reqType == "check" {
		for queryResult.Next() {
			if err := queryResult.Scan(&scanResult[0], &scanResult[1], &scanResult[2]); err != nil {
				log.Fatal(err)
			}
		}
	}

	return scanResult
}

func insertData(w http.ResponseWriter, r *http.Request) {
	var data = unmarshalData(w, r)
	insertSQL := fmt.Sprintf(`
			INSERT INTO reg_data (email, username, password)
			VALUES ('%v', '%v', '%v');`, data.Email, data.Username, data.Password)
	doQuery(insertSQL, "insert")
}

func checkData(w http.ResponseWriter, r *http.Request) {
	var data = unmarshalData(w, r)

	emailCheck := fmt.Sprintf(`select * from user_data where email="%v"`, data.Email)
	userCheck := fmt.Sprintf(`select * from user_data where username="%v"`, data.Username)
	passwordCheck := fmt.Sprintf(`select * from user_data where username="%v"`, data.Password)

	emailResult := doQuery(emailCheck, "check")[0]
	userResult := doQuery(userCheck, "check")[1]
	passwordResult := doQuery(passwordCheck, "check")[2]

}

func main() {
	//Handling requests and endpoints
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/checkData", checkData).Methods("POST")
	api.HandleFunc("/insertData", insertData).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	//cors related stuff
	handler := c.Handler(r)

	//Listening for requests
	log.Fatal(http.ListenAndServe(":8081", handler))
}
