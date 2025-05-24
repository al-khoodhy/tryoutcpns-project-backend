package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
	"tryoutcpns-project-backend/utils"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if user.Email == "" || user.Phone == "" || user.Password == "" {
		http.Error(w, "Email, phone, and password are required", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Cek apakah email atau phone sudah ada
	var existingUser models.User
	if err := config.DB.Where("email = ? OR phone = ?", user.Email, user.Phone).First(&existingUser).Error; err == nil {
		http.Error(w, "Email or phone already exists", http.StatusConflict)
		return
	}

	// Simpan user baru terlebih dahulu untuk mendapatkan ID
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Generate referral ID dan update user
	user.ReferralID = utils.GenerateReferralID(user.ID, user.Email)
	if err := config.DB.Model(&user).Update("referral_id", user.ReferralID).Error; err != nil {
		http.Error(w, "Failed to update referral ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "User registered successfully",
		"referral_id": user.ReferralID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
