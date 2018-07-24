package main

import (
	"log"
	"net/http"

	"github.com/dmrajkarthick/TweetStack/config"
	"github.com/dmrajkarthick/TweetStack/dbo"
	"github.com/gorilla/mux"
	"github.com/dmrajkarthick/TweetStack/handler"
)

var conf config.Config
var dboper dbo.DBOperations

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
	r.HandleFunc("/questions", handler.GetAllQuestions).Methods("GET")
	r.HandleFunc("/questions", handler.CreateQuestion).Methods("POST")
	r.HandleFunc("/questions", handler.UpdateQuestion).Methods("PUT")
	r.HandleFunc("/questions", handler.DeleteQuestion).Methods("DELETE")
	r.HandleFunc("/questions/{id}", handler.FindQuestionById).Methods("GET")
	r.HandleFunc("/answers/question/{questionId}", handler.GetAllAnswers).Methods("GET")
	r.HandleFunc("/answers/{questionId}", handler.AddAnswer).Methods("POST")
	r.HandleFunc("/answers", handler.UpdateAnswer).Methods("PUT")
	r.HandleFunc("/answers", handler.DeleteAnswer).Methods("DELETE")
	r.HandleFunc("/answer/{id}", handler.FindAnswerById).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
