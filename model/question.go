package model

import "github.com/globalsign/mgo/bson"

// Represents a question, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Question struct {
	Id          bson.ObjectId `bson:"id,omitempty" 	json:"moid"`
	Question    string        `bson:"question" json:"question"`
	Tags        []string      `bson:"tags" json:"tags"`
	Description string        `bson:"description" json:"description"`
}
