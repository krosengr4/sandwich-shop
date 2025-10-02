package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadEnv(filename string) error {
	// Open the file passed in
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open the .env file: %w", err)
	}

	// Create new scanner
	scanner := bufio.NewScanner(file)
	// Scan each line of file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split on the "=" Skip if anything other than 2 strings returned
		lineParts := strings.SplitN(line, "=", 2)
		if len(lineParts) != 2 {
			continue
		}

		// Declare the key and the value
		key := strings.TrimSpace(lineParts[0])
		value := strings.TrimSpace(lineParts[1])

		// Set the env variables (using os.Setenv)
		os.Setenv(key, value)
	}

	return scanner.Err()
}

// Retrieves db config from env variables
func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
