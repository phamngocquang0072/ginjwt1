package initializers

import (
	"github.com/joho/godotenv"
)

func GetEnvVariable() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}