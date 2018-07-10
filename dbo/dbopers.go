package dbo

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBOperations struct {
	Server   string
	Database string
}

var db *mgo.Database

// Establish a connection to database
func (m *DBOperations) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *DBOperations) FindAll(collection string) ([]bson.M, error) {
	var obj []bson.M
	err := db.C(collection).Find(bson.M{}).All(&obj)
	return obj, err
}

// Find a collection by its id
func (m *DBOperations) FindOne(collection string, id string) (bson.M, error) {
	var obj bson.M
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&obj)
	//err := db.C(collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&obj)
	return obj, err
}

// Insert a collection into database
func (m *DBOperations) Insert(collection string, obj interface{}) error {
	err := db.C(collection).Insert(&obj)
	return err
}

// Delete an existing collection
func (m *DBOperations) Delete(collection string, obj interface{}) error {
	err := db.C(collection).Remove(&obj)
	return err
}

// Update an existing collection
func (m *DBOperations) Update(collection string, id bson.ObjectId, obj interface{}) error {
	err := db.C(collection).UpdateId(id , &obj)
	return err
}
