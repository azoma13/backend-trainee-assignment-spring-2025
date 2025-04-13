package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	APIPort      string
	SecretJWTKey []byte
)

func Environment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error load .env: %w", err)
	}

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	APIPort = os.Getenv("API_PORT")
	SecretJWTKey = []byte(os.Getenv("SECRET_KEY"))

	return nil
}
