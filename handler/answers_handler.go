package handler

import (
	"net/http"
	"github.com/dmrajkarthick/TweetStack/model"
	"github.com/dmrajkarthick/TweetStack/utils"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

var dboper_answers dbo.DBOperations

func GetAllAnswers(w http.ResponseWriter, r *http.Request){
	var answers []model.Answer

	res, err := dboper_answers.FindAll(utils.COLLECTION_ANSWERS)
	if err != nil{
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = mapstructure.Decode(res, &answers)
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

