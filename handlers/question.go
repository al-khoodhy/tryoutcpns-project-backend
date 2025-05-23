package handlers

import (
    "tryoutcpns-project-backend/models"
    "tryoutcpns-project-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetQuestion(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var question models.Question
    if err := config.DB.First(&question, id).Error; err != nil {
        http.Error(w, "Question not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(question)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
    var questions []models.Question
    config.DB.Find(&questions)
    json.NewEncoder(w).Encode(questions)
}