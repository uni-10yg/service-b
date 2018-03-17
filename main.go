package main

import (
	"log"
	"net/http"
	//"encoding/json"
	"github.com/gorilla/mux"
)

func postCredentials(resp_writer http.ResponseWriter, request *http.Request) {

}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", postCredentials).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}