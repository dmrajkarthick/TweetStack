package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dmrajkarthick/TweetStack/config"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"github.com/dmrajkarthick/TweetStack/model"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var conf config.Config
var dboper dbo.DBOperations

// GET list of questions
func AllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := dboper.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, questions)
}

// GET a question by its ID
func FindQuestionById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	question, err := dboper.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Question ID")
		return
	}
	respondWithJson(w, http.StatusOK, question)
}

// POST a new question
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	question.ID = bson.NewObjectId()
	if err := dboper.Insert(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, question)
}

// PUT update an existing question
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper.Update(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing question
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dboper.Delete(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	conf.Read()

	dboper.Server = conf.Server
	dboper.Database = conf.Database
	dboper.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/questions", AllQuestions).Methods("GET")
	r.HandleFunc("/questions", CreateQuestion).Methods("POST")
	r.HandleFunc("/questions", UpdateQuestion).Methods("PUT")
	r.HandleFunc("/questions", DeleteQuestion).Methods("DELETE")
	r.HandleFunc("/questions/{id}", FindQuestionById).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
