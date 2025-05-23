package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadEnv memuat variabel lingkungan dari file .env
func LoadEnv() error {
	// Cek apakah file .env ada
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("⚠️ File .env tidak ditemukan. Pastikan file tersebut ada di root proyek.")
		return nil // Tidak error jika file tidak ada (opsional)
	}

	file, err := os.Open(".env")
	if err != nil {
		return fmt.Errorf("gagal membuka file .env: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Lewati baris kosong atau komentar
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Abaikan baris yang tidak valid
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set environment variable
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("gagal membaca file .env: %w", err)
	}

	return nil
}
