package handler

import (
	"fmt"
	"net/http"
	"github.com/dmrajkarthick/TweetStack/model"
	"github.com/dmrajkarthick/TweetStack/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"encoding/json"
)


var dboper_questions dbo.DBOperations

// GET list of questions
func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	var questions []model.Question
	res, err := dboper_questions.FindAll(utils.COLLECTION_QUESTIONS)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(res)
	jsonData, err := json.Marshal(res)
	json.Unmarshal(jsonData, &questions)
	
	utils.RespondWithJson(w, http.StatusOK, questions)
}

// GET a question by its ID
func FindQuestionById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var question model.Question
	res, err := dboper_questions.FindOne(utils.COLLECTION_QUESTIONS, params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Question ID")
		return
	}
	jsonData, err := json.Marshal(res)

	json.Unmarshal(jsonData, &question)
	utils.RespondWithJson(w, http.StatusOK, question)
}

// POST a new question
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	question.ID = bson.NewObjectId()

	if err := dboper_questions.Insert(utils.COLLECTION_QUESTIONS, question); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, question)
}

// PUT update an existing question
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper_questions.Update(utils.COLLECTION_QUESTIONS, question.ID, question); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "successfully updated"})
}

// DELETE an existing question
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper_questions.Delete(utils.COLLECTION_QUESTIONS, question); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "successfully deleted"})
}

