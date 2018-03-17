package main

import (
	"log"
	"net/http"
	"math/rand"
	"time"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

var amqp_client IamqpClient

type SaltedHash struct {
	Hash	string	`json:"hash"`
	Salt	int		`json:"salt"`
}

type Phrase struct {
	Value	string `json:"phrase"`
}

func randomFromRange(low, high int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(high-low) + low
}

func hashAndSalt(phrase string) SaltedHash {
	salt := randomFromRange(bcrypt.MinCost, bcrypt.DefaultCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(phrase), salt)
	if err != nil {
		panic(err)
	}
	salted_hash := SaltedHash{ string(hash), salt}
	return salted_hash
}

func postPhrase(resp_writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	var phrase Phrase
	if err := json.Unmarshal(body, &phrase); err != nil {
		resp_writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp_writer.WriteHeader(422)
		if err := json.NewEncoder(resp_writer).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("got: %s", phrase)
	salted_hash_b, err := json.Marshal(hashAndSalt(phrase.Value))
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated: %s", salted_hash_b)

	//amqp_client.sendMsg(salted_hash_b, "comm_queue")

	resp_writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp_writer.WriteHeader(http.StatusCreated)
	resp_writer.Write(salted_hash_b)
}

func main() {
	amqp_client = &AmqpClient{}
	//amqp_client.connectToBroker("")
	router := mux.NewRouter()
	router.HandleFunc("/", postPhrase).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}