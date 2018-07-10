package handler

import (
	"net/http"
	"github.com/dmrajkarthick/TweetStack/model"
	"github.com/dmrajkarthick/TweetStack/utils"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

var dboper_answers dbo.DBOperations

func GetAllAnswers(w http.ResponseWriter, r *http.Request){
	var answers []model.Answer

	res, err := dboper_answers.FindAll(utils.COLLECTION_ANSWERS)
	if err != nil{
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	jsonData, err := json.Marshal(res)
	json.Unmarshal(jsonData, &answers)
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

	var answer model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err!=nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	answer.ID = bson.NewObjectId()
	if err := dboper_answers.Insert(utils.COLLECTION_ANSWERS, answer); err!=nil {
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









