package config

import (
	"github.com/joho/godotenv"
	"os"
)

func NewConfig() HlsProvider {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		panic("SERVICE_PORT is not set")
	}

	return HlsProvider{ServicePort: ServicePort(port)}
}
