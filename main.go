package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter()
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}

func main() {
	handleRequests()
}