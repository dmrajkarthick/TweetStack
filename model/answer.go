package model

import "github.com/globalsign/mgo/bson"

// Represents a answer, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Answer struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Answer  string        `bson:"question" json:"question"`
	Tags    []string      `bson:"tags" json:"tags"`
	Upvotes int           `bson:"upvotes" json:"Upvotes"`
}
