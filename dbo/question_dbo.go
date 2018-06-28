package dao

import (
	"log"

	. "github.com/dmrajkarthick/TweetStack/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QuestionDBO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Questions"
)

// Establish a connection to database
func (m *QuestionDBO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of Questions
func (m *QuestionDBO) FindAll() ([]Question, error) {
	var Questions []Question
	err := db.C(COLLECTION).Find(bson.M{}).All(&Questions)
	return Questions, err
}

// Find a Question by its id
func (m *QuestionDBO) FindById(id string) (Question, error) {
	var Question Question
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Question)
	return Question, err
}

// Insert a Question into database
func (m *QuestionDBO) Insert(Question Question) error {
	err := db.C(COLLECTION).Insert(&Question)
	return err
}

// Delete an existing Question
func (m *QuestionDBO) Delete(Question Question) error {
	err := db.C(COLLECTION).Remove(&Question)
	return err
}

// Update an existing Question
func (m *QuestionDBO) Update(Question Question) error {
	err := db.C(COLLECTION).UpdateId(Question.ID, &Question)
	return err
}
