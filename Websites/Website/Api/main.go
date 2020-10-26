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

func validation() {

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

		//Input data (User data)
		fmt.Println("Email =", data.Email)
		fmt.Println("Username =", data.Username)
		fmt.Println("Password =", data.Password, "\n", "")

		//Database connection.
		const (
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "shashi3969"
			dbname   = "Glade"
		)

		//Commands for the database.
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

		//Checking
		emailValidateSQL := fmt.Sprintf(`select * from reg_data where email='%v'`, data.Email)
		userValidateSQL := fmt.Sprintf(`select * from reg_data where user='%v'`, data.Email)
		emailExec, err := db.Query(emailValidateSQL)
		for emailExec.Next() {
			var (
				Email    string
				Username string
			)
			if err := emailExec.Scan(&Email, &Username); err != nil {
				log.Fatal(err)
			}
			fmt.Println(Email, Username)
		}
		if err != nil {
			fmt.Println("Validation error : ", err)
		}

		/*	insertSQL := fmt.Sprintf(`
			INSERT INTO reg_data (email, username, password, userkey)
			VALUES ('%v', '%v', '%v');`, data.Email, data.Username, data.Password)
				_, err = db.Exec(insertSQL)
				if err != nil {
					fmt.Println("Inserting values error : ", insertSQL)
				}*/

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
