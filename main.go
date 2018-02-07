package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/persons", GetAllPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/persons", CreatePerson).Methods("POST")
	router.HandleFunc("/persons/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
