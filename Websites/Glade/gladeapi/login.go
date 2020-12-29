package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	unmarshalData(w, r)
	var data Data

	userQuery := fmt.Sprintf(`select * from user_data where username="%v"`, data.Username)
	passwordQuery := fmt.Sprintf(`select * from user_data where username="%v"`, data.Password)

	userResult := queryDb(userQuery)[1]
	passwordResult := queryDb(passwordQuery)[2]

	if userResult != nil && passwordResult != nil {
		respond(w, r, "logMatch12")
	} else if userResult != nil && passwordResult == nil {
		respond(w, r, "logMatch1")
	} else if userResult == nil && passwordResult != nil {
		respond(w, r, "logMatch2")
	} else if userResult == nil && password == nil {
		respond(w, r, "logMatch0")
	}
}
