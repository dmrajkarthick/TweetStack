package model

type DBRef struct {
	Collection string      `bson:"collection" json:"collection"`
	Id         interface{} `bson:"_id" json:"_id"`
	Database   string      `bson:"database,omitempty" json:"database"`
}
