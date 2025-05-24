package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateReferralID generates a user-friendly unique referral code based on user ID, email, and timestamp
func GenerateReferralID(userID uint, email string) string {
	email = strings.ToLower(email)
	hash := sha1.New()
	hash.Write([]byte(email))
	emailHash := hex.EncodeToString(hash.Sum(nil))[:3] // ambil 3 huruf dari hash email

	timePart := time.Now().UnixNano()
	randPart := rand.Intn(999)

	return fmt.Sprintf("CPNS-%03d%s%03d", userID, strings.ToUpper(emailHash), randPart+int(timePart%1000))
}
