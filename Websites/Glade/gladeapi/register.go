package main

import "net/http"

func register(w http.ResponseWriter, r *http.Request) {
	data := unmarshalData(w, r)
}
