package env

import (
	"os"

	"github.com/joho/godotenv"
)

var DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME string

func Getenv(path string) error {
	// 環境変数のloading
	if err := godotenv.Load(path); err != nil {
		return err
	}
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	return nil
}