package service

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type env struct {
	Mgourl string
	Mgodb  string
	Port   string
}

var ENV_CONFIG env

func InitEnv() {
	absPath, _ := filepath.Abs("../.env")
	err := godotenv.Load(absPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ENV_CONFIG.Mgourl = os.Getenv("MONGO_URL")
	ENV_CONFIG.Mgodb = os.Getenv("MGO_DB")

	ENV_CONFIG.Port = os.Getenv("BASE_PORT")
}
