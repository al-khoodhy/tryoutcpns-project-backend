package handlers

import (
    "cpns-backend/models"
    "cpns-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetUserAnswer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var answer models.UserAnswer
    if err := config.DB.First(&answer, id).Error; err != nil {
        http.Error(w, "Answer not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(answer)
}

func GetAllUserAnswers(w http.ResponseWriter, r *http.Request) {
    var answers []models.UserAnswer
    config.DB.Find(&answers)
    json.NewEncoder(w).Encode(answers)
}