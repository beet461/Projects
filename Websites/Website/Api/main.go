package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//Data ..
type Data struct {
	Email    string
	Username string
	Password string
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func insertData(data Data) {
	insertSQL := fmt.Sprintf(`
			INSERT INTO reg_data (email, username, password)
			VALUES ('%v', '%v', '%v');`, data.Email, data.Username, data.Password)
	dataQuery(insertSQL, "Inserting Data")
}

func dataQuery(query, dataType string) string {
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

	//Query the statement in the database
	queryResult, err := db.Query(query)
	if err != nil {
		fmt.Println("Validation error : ", err)
	}

	//Check for any existing values that match the input
	var scanResult[3]string
	for queryResult.Next() {
		if err := queryResult.Scan(&scanResult[0]); err != nil {
			log.Fatal(err)
		}
		fmt.Println(dataType, scanResult[0], scanResult[1], scanResult[2])
	}
	return scanResult[0], 
}

func register(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
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
		var data Data
		json.Unmarshal([]byte(body), &data)

		//Input data
		fmt.Println("Email =", data.Email)
		fmt.Println("Username =", data.Username)
		fmt.Println("Password =", data.Password, "\n", "")

		//Sql statements
		emailValidateSQL := fmt.Sprintf(`select * from reg_data where email='%v'`, data.Email)
		userValidateSQL := fmt.Sprintf(`select * from reg_data where user='%v'`, data.Username)
		passwordValidateSQL := fmt.Sprintf(`select * from reg_data where password='%v'`, data.Password)

		//Query is run
		emailResult := dataQuery(emailValidateSQL, "Email : ")
		userResult := dataQuery(userValidateSQL, "Username : ")
		passwordResult := dataQuery(passwordValidateSQL, "Password : ")

		if emailResult == "" && userResult == "" && passwordResult == "" {
			insertData(data)
		} else {
			fmt.Println("One of the values is the same")
		}

	}
}

func main() {
	//Handling requests and endpoints
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/register", register).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	//cors related stuff
	handler := c.Handler(r)

	//Listening for requests
	log.Fatal(http.ListenAndServe(":8081", handler))
}
