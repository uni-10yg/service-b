package main

import (
	"log"
	"net/http"
	"math/rand"
	//"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"time"
	"encoding/json"
)

type SaltedHash struct {
	Hash	string	`json:"hash"`
	Salt	int		`json:"salt"`
}

func randomFromRange(low, high int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(high-low) + low
}

func hashAndSalt(phrase string) SaltedHash {
	salt := randomFromRange(bcrypt.MinCost, bcrypt.MaxCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(phrase), salt)
	if err != nil {
		log.Println(err)
	}
	salted_hash := SaltedHash{ string(hash), salt}
	return salted_hash
}

func postPhrase(resp_writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	salted_hash := hashAndSalt(params["phrase"])
	//bytes, err := json.Marshal(salted_hash)
	//if err != nil {
	//	log.Println(err)
	//}
	log.Println(salted_hash)
	json.NewEncoder(resp_writer).Encode(salted_hash)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", postPhrase).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}

func main() {
	handleRequests()
}