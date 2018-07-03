package dbo

import (
	"log"

	"github.com/dmrajkarthick/TweetStack/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBOperations struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "questions"
)

// Establish a connection to database
func (m *DBOperations) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of Questions
func (m *DBOperations) FindAll() ([]model.Question, error) {
	var Questions []model.Question
	err := db.C(COLLECTION).Find(bson.M{}).All(&Questions)
	return Questions, err
}

// Find a Question by its id
func (m *DBOperations) FindById(id string) (model.Question, error) {
	var Question model.Question
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Question)
	return Question, err
}

// Insert a Question into database
func (m *DBOperations) Insert(Question model.Question) error {
	err := db.C(COLLECTION).Insert(&Question)
	return err
}

// Delete an existing Question
func (m *DBOperations) Delete(Question model.Question) error {
	err := db.C(COLLECTION).Remove(&Question)
	return err
}

// Update an existing Question
func (m *DBOperations) Update(Question model.Question) error {
	err := db.C(COLLECTION).UpdateId(Question.ID, &Question)
	return err
}
