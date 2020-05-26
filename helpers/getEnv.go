package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {log.Fatal(err)}
	return os.Getenv(key)
}
