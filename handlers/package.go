package handlers

import (
    "tryoutcpns-project-backend/models"
    "tryoutcpns-project-backend/config"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
)

func GetPackage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var packageData models.Package
    if err := config.DB.First(&packageData, id).Error; err != nil {
        http.Error(w, "Package not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(packageData)
}

func GetAllPackages(w http.ResponseWriter, r *http.Request) {
    var packages []models.Package
    config.DB.Find(&packages)
    json.NewEncoder(w).Encode(packages)
}