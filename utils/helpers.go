package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword menghash password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPasswordHash memverifikasi password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWTToken generates a JWT token with user ID and expiration
func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Ganti dengan secret key dari .env
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken validates the JWT token and returns the user ID
func ValidateJWTToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// Ganti dengan secret key dari .env
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(float64)
		if !ok {
			return 0, fmt.Errorf("invalid user ID in token")
		}
		return uint(userID), nil
	}

	return 0, fmt.Errorf("invalid token")
}

// GenerateReferralCode generates a unique UUID for affiliate
func GenerateReferralCode() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

// FormatDate converts time.Time to string format "YYYY-MM-DD HH:mm:ss"
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// IsEmpty checks if a string is empty or only contains whitespace
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNil checks if a pointer is nil
func IsNil(i interface{}) bool {
	return i == nil
}

// NullString handles SQL NULL string values
func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// NullInt64 handles SQL NULL int64 values
func NullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

// NullBool handles SQL NULL bool values
func NullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}
