package config

import (
	//"log"
	"os"
	//"github.com/joho/godotenv"
)

var (
	DB_user     = ""
	DB_password = ""
	DB_port     = ""
	DB_host     = ""
	DB_name     = ""
	API_port = ""
)

func LoadVar() {

	DB_user = os.Getenv("POSTGRES_USER")
	DB_password = os.Getenv("POSTGRES_PASSWORD")
	DB_port = os.Getenv("POSTGRES_PORT")
	DB_host = os.Getenv("POSTGRES_HOST")
	DB_name = os.Getenv("POSTGRES_NAME")
	API_port = os.Getenv("API_PORT")
}
