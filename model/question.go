package model

import "gopkg.in/mgo.v2/bson"

// Represents a question, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Question struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Question    string        `bson:"question" json:"question"`
	Tags        []string      `bson:"tags" json:"tags"`
	Description string        `bson:"description" json:"description"`
}
