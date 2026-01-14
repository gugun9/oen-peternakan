package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Gagal load file .env")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}