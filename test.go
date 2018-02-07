package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session, err = mgo.Dial("localhost:27017")
var c = session.DB("go-test").C("persons")

//GetAllPersons endpoint
func GetAllPersons(w http.ResponseWriter, req *http.Request) {
	var persons []Person
	err = c.Find(bson.M{}).All(&persons)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

//GetPerson endpoint
func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&person)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)

}

//CreatePerson endpoint
func CreatePerson(w http.ResponseWriter, req *http.Request) {

	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	newID := bson.NewObjectId()

	_, err := c.UpsertId(newID, person)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Find(bson.M{"_id": newID}).One(&person)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)

}

//DeletePerson endpoint
func DeletePerson(w http.ResponseWriter, req *http.Request) {

}
