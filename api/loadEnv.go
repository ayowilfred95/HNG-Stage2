package api

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// check if production

	production := os.Getenv("PRO")

	if production != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
