package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

//GetAllPersons endpoint
func GetAllPersons(c *mgo.Collection) func(http.ResponseWriter, *http.Request) {
	if c == nil {
		panic("nil MongoDB session!")
	}

	return func(w http.ResponseWriter, req *http.Request) {
		var persons []Person
		err := c.Find(bson.M{}).All(&persons)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}

//GetPerson endpoint
func GetPerson(c *mgo.Collection) func(http.ResponseWriter, *http.Request) {
	if c == nil {
		panic("nil MongoDB session!")
	}
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var person Person
		err := c.Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&person)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(person)
	}
}

//CreatePerson endpoint
func CreatePerson(c *mgo.Collection) func(http.ResponseWriter, *http.Request) {
	if c == nil {
		panic("nil MongoDB session!")
	}
	return func(w http.ResponseWriter, req *http.Request) {
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
}
