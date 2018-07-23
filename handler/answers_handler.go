package handler

import (
	"net/http"
	"github.com/dmrajkarthick/TweetStack/model"
	"github.com/dmrajkarthick/TweetStack/utils"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var dboper_answers dbo.DBOperations

func GetAllAnswers(w http.ResponseWriter, r *http.Request){
	var answers []model.Answer

	params := mux.Vars(r)
	questionId := params["questionId"]
	var qn model.Question
	res, err := dboper_questions.FindOne(utils.COLLECTION_QUESTIONS, questionId)
	if err!=nil{
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	jsonData, err := json.Marshal(res)
	json.Unmarshal(jsonData, &qn)
	
	answers, err = qn.GetRelAnswers()
	if err != nil{
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	utils.RespondWithJson(w, http.StatusOK, answers)
}

//Find an answer using its ID
func FindAnswerById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	answer, err := dboper_answers.FindOne(utils.COLLECTION_ANSWERS, params["id"])
	if err != nil{
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Answer Id")
	}
	utils.RespondWithJson(w, http.StatusOK, answer)
}

func AddAnswer(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	params := mux.Vars(r)
	questionId := params["questionId"]

	if questionId == ""{
		utils.RespondWithError(w, http.StatusInternalServerError, "Please provide details of the question for this answer")
	}

	var answer model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err!=nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//Adding question details to answer
	answer.ID = bson.NewObjectId()
	qn := mgo.DBRef{
		Collection: utils.COLLECTION_QUESTIONS,
		Id: questionId,
		Database:"",
	}
	answer.QuestionId = qn

	if err := dboper_answers.Insert(utils.COLLECTION_ANSWERS, answer); err!=nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//Adding answer details to question.
	var question model.Question
	res, err := dboper_questions.FindOne(utils.COLLECTION_QUESTIONS, questionId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Question ID")
		return
	}
	jsonData, err := json.Marshal(res)

	json.Unmarshal(jsonData, &question)
	question.SetRelAnswer(answer)

	if err := dboper_questions.Update(utils.COLLECTION_QUESTIONS, question.ID, question); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, answer)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var answer model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil{
		utils.RespondWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper_answers.Update(utils.COLLECTION_ANSWERS, answer.ID, answer); err != nil{
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "successfully updated"})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var answer model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper_answers.Delete(utils.COLLECTION_ANSWERS, answer); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "successfully deleted"})
}









