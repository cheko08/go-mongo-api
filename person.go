package main

import "gopkg.in/mgo.v2/bson"

// Person struct
type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName string        `bson:"firstName" json:"firstName,omitempty"`
	LastName  string        `bson:"lastName" json:"lastName,omitempty"`
	Address   *Address      `json:"address,omitempty"`
}
