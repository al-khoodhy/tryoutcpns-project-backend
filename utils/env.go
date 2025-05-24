package utils

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv loads environment variables from a .env file manually (without external lib)
func LoadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Lewati baris kosong atau komentar
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Pecah berdasarkan tanda '=' hanya sekali
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Hilangkan kutipan di value jika ada
		value = strings.Trim(value, `"'`)

		// Set ke environment
		_ = os.Setenv(key, value)
	}

	// Cek error scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
