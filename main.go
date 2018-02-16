package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	router := mux.NewRouter()

	var session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	Collection := session.DB("go-test").C("persons")

	router.HandleFunc("/persons", GetAllPersons(Collection)).Methods("GET")
	router.HandleFunc("/persons/{id}", GetPerson(Collection)).Methods("GET")
	router.HandleFunc("/persons", CreatePerson(Collection)).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
