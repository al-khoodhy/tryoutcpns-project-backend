package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"

	"github.com/gorilla/mux"
)

func GetVoucher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var voucher models.Voucher
	if err := config.DB.First(&voucher, id).Error; err != nil {
		http.Error(w, "Voucher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(voucher)
}

func GetAllVouchers(w http.ResponseWriter, r *http.Request) {
	var vouchers []models.Voucher
	config.DB.Find(&vouchers)
	json.NewEncoder(w).Encode(vouchers)
}
