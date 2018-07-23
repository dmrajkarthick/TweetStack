package model

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/dmrajkarthick/TweetStack/utils"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"encoding/json"
	"gopkg.in/mgo.v2"
)

// Represents a question, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Question struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Question    string        `json:"question" bson:"question"`
	Tags        []string      `json:"tags" bson:"tags"`
	Upvotes 	int           `json:"upvotes" bson:"upvotes"`
	Description string        `json:"description" bson:"description"`

	//List of Answer IDs that are associated to a question.
	AnswerIds   []DBRef   `json:"answerIds" bson:"answerIds"`
}

var dboper_questions dbo.DBOperations

//TODO: Populate this.
type QuestionIf interface {
	GetID() bson.ObjectId
	SetID(ID bson.ObjectId)

	GetQuestion() string
	SetQuestion(Question string)

	GetTags() []string
	SetTags(Tags []string)

	GetUpvotes() int
	SetUpvotes(int)

	GetDescription() string
	SetDescription(string)

	GetRelAnswerIds() []mgo.DBRef
	SetRelAnswerIds([]mgo.DBRef)

}


func (question *Question) GetID() bson.ObjectId{
	return question.ID
}

func (question *Question) SetID(ID bson.ObjectId){

}

func (question *Question) GetQuestion() string{
	return question.Question
}

func (question *Question) SetQuestion(Question string){

}

func (question *Question) GetTags() []string{
	return question.Tags
}

func (question *Question) SetTags(tags []string){
	question.Tags = tags
}

func (question *Question) GetUpvotes() int{
	return 0
}

func (question *Question) SetUpvotes(int){

}

func (question *Question) GetDescription() string{
	return question.Description
}

func (question *Question) SetDescription(description string){
	question.Description = description
}


func (question *Question) GetRelAnswerIds() []DBRef{
	return question.AnswerIds
}

func (question *Question) SetRelAnswerIds([]DBRef){

}

func (question *Question) GetRelAnswers() ([]Answer, error){
	// Get answers DBRefs and retrive data from DB.
	var answers []Answer
	var answer Answer
	for _, v := range question.AnswerIds {
		res, err := dboper_questions.FindOne(v.Collection, v.Id.(string))
		if err != nil {
			return nil, err
		}
		jsonData, err := json.Marshal(res)

		json.Unmarshal(jsonData, &answer)
		answers = append(answers, answer)
	}
	return answers, nil
}

func (question *Question) SetRelAnswers(answer []Answer) error{
	for _, v := range answer{
		ans := DBRef{
			Collection: utils.COLLECTION_ANSWERS,
			Id: v.ID,
			Database:"",
		}
		question.AnswerIds = append(question.AnswerIds, ans)
	}

	// perform patch after adding answer details.
	if err := dboper_questions.Update(utils.COLLECTION_QUESTIONS, question.ID, question); err != nil {
		return err
	}

	return nil
}

func (question *Question) SetRelAnswer(answer Answer){
	ans := DBRef{
		Collection: utils.COLLECTION_ANSWERS,
		Id: answer.ID,
		Database:"",
	}
	question.AnswerIds = append(question.AnswerIds, ans)
}





