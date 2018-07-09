package model

import "gopkg.in/mgo.v2/bson"

// Represents a answer, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Answer struct {
	ID     bson.ObjectId  `bson:"_id,omitempty" json:"_id"`
	Answer  string        `bson:"answer" json:"answer"`
	Tags    []string      `bson:"tags" json:"tags"`
	Upvotes int           `bson:"upvotes" json:"upvotes"`
}
