package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

// Represents a answer, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Answer struct {
	ID     bson.ObjectId  `bson:"_id,omitempty" json:"_id"`
	Answer  string        `bson:"answer" json:"answer"`
	Tags    []string      `bson:"tags" json:"tags"`
	Upvotes int           `bson:"upvotes" json:"upvotes"`
	Description string    `bson:"description" json:"description"`

	//Question Id to which an answer is related. 
	QuestionId   DBRef   `bson:"questionId" json:"questionId"`
}

type AnswerIf interface {
	GetID() bson.ObjectId
	SetID(ID bson.ObjectId)

	GetAnswer() string
	SetAnswer(Question string)

	GetTags() []string
	SetTags(Tags []string)

	GetUpvotes() int
	SetUpvotes(int)

	GetDescription() string
	SetDescription(string)


	GetRelQuestionId() mgo.DBRef
	SetRelQuestionId(mgo.DBRef)
}


func (answer *Answer) GetID() bson.ObjectId{
	return answer.ID
}

func (answer *Answer) SetID(ID bson.ObjectId){

}

func (answer *Answer) GetAnswer() string{
	return answer.Answer
}

func (answer *Answer) SetAnswer(Question string){

}

func (answer *Answer) GetTags() []string{
	return answer.Tags
}

func (answer *Answer) SetTags(Tags []string){

}

func (answer *Answer) GetUpvotes() int{
	return answer.Upvotes
}

func (answer *Answer) SetUpvotes(int){

}

func (answer *Answer) GetDescription() string{
	return answer.Description
}

func (answer *Answer) SetDescription(description string){
	answer.Description = description
}

func (answer *Answer) GetRelQuestionId() DBRef{
	return answer.QuestionId
}

func (answer *Answer) SetRelQuestionId(DBRef){

}

