package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Articles{
		Article{Title: "Titlele", Desc: "Le Describetion", Content: "Articlecle Contentent"}
		Article{Title: "Titlele 2", Desc: "Le Describetion 2", Content: "Articlecle Contentent 2"}
	}
	handleRequests()
}
