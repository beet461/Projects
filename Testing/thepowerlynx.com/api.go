package com

import (
	"net/http"

	"github.com/gorilla/mux"
)

func web() {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/website", web)
	http.ListenAndServe("www.thepowerlynx.com", r)
}
