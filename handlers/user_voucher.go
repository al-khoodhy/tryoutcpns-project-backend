package handlers

import (
    "cpns-backend/models"
    "cpns-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func ApplyVoucher(w http.ResponseWriter, r *http.Request) {
    var voucherData struct {
        UserID   uint   `json:"user_id"`
        VoucherID uint   `json:"voucher_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&voucherData); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validasi input
    if voucherData.UserID == 0 || voucherData.VoucherID == 0 {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    // Cek apakah user ada
    var user models.User
    if err := config.DB.First(&user, voucherData.UserID).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Cek apakah voucher ada
    var voucher models.Voucher
    if err := config.DB.First(&voucher, voucherData.VoucherID).Error; err != nil {
        http.Error(w, "Voucher not found", http.StatusNotFound)
        return
    }

    // Cek apakah voucher sudah digunakan
    var usedVoucher models.UserVoucher
    if err := config.DB.Where("user_id = ? AND voucher_id = ?", voucherData.UserID, voucherData.VoucherID).First(&usedVoucher).Error; err == nil {
        http.Error(w, "Voucher already used", http.StatusConflict)
        return
    }

    // Tambahkan voucher ke user
    userVoucher := models.UserVoucher{
        UserID:   voucherData.UserID,
        VoucherID: voucherData.VoucherID,
    }

    if err := config.DB.Create(&userVoucher).Error; err != nil {
        http.Error(w, "Failed to apply voucher", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Voucher applied successfully"})
}