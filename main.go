package main

import (
	"log"
	"net/http"
	"math/rand"
	//"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func randomFromRange(low, high int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(high-low) + low
}

func hashAndSalt(phrase []byte) (string, int) {
	var salt = randomFromRange(bcrypt.MinCost, bcrypt.MaxCost)
	hash, err := bcrypt.GenerateFromPassword(phrase, salt)
	if err != nil {
		log.Println(err)
	}
	return string(hash), salt
}

func postPhrase(resp_writer http.ResponseWriter, request *http.Request) {
	//params := mux.Vars(request)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", postPhrase).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}

func main() {
	handleRequests()
}