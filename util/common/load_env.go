package common

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file")
	}

	return nil
}
