package handlers

import (
    "tryoutcpns-project-backend/models"
    "tryoutcpns-project-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetAffiliate(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var affiliate models.Affiliate
    if err := config.DB.First(&affiliate, id).Error; err != nil {
        http.Error(w, "Affiliate not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(affiliate)
}

func GetAllAffiliates(w http.ResponseWriter, r *http.Request) {
    var affiliates []models.Affiliate
    config.DB.Find(&affiliates)
    json.NewEncoder(w).Encode(affiliates)
}