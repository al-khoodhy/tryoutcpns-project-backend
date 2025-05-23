package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
)

func GetMateri(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var materi models.Materi
	if err := config.DB.First(&materi, id).Error; err != nil {
		http.Error(w, "Materi not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(materi)
}

func GetAllMateri(w http.ResponseWriter, r *http.Request) {
	var materi []models.Materi
	config.DB.Find(&materi)
	json.NewEncoder(w).Encode(materi)
}
