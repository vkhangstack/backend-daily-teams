package repository

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/vkhangstack/dlt/internal/config"
)

func LoadAPIConfig() (*config.APIConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	apiKey := os.Getenv("API_KEY")

	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret not found")
	}

	return &config.APIConfig{
		JWTSecret: jwtSecret,
		APIKey:    apiKey,
	}, nil
}
