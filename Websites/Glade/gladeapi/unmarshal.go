package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func unmarshalData(w http.ResponseWriter, r *http.Request) {
	var data Data
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
		json.Unmarshal([]byte(body), &data)
	}
}
