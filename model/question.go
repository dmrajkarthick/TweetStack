package model

import "gopkg.in/mgo.v2/bson"

// Represents a question, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Question struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Question    string        `json:"question" bson:"question"`
	Tags        []string      `json:"tags" bson:"tags"`
	Description string        `json:"description" bson:"description"`
}
