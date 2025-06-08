package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading %s .env file\n", key)
	}
	return os.Getenv(key)
}

// Return casted to integer config value (default = 0)
func ConfigInt(key string) int {
	i, err := strconv.Atoi(Config(key))
	if err != nil {
		return 0
	}
	return i
}