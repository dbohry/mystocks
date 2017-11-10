package models

import "gopkg.in/mgo.v2/bson"

// User entity
//
type User struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email string `json:"email,omitempty"`
	User string `json:"user,omitempty"`
	Pass string `json:"pass,omitempty"`
}