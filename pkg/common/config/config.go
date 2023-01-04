package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT            string
	DB_URL          string
	API_SECRET      string
	TOKEN_LIFE_SPAN int
}

func InitConfig() error {
	var err error

	wd, _ := os.Getwd()
	fmt.Println("wd:")
	fmt.Println(wd);
	fmt.Println("")

	if err = godotenv.Load("./config.env"); err != nil {
		log.Print("No .env file found")
	}

	return err
}

func GetConfig() *Config {
	return &Config{
		PORT:            getEnvAsString("PORT", ""),
		DB_URL:          getEnvAsString("DB_URL", ""),
		API_SECRET:      getEnvAsString("API_SECRET", ""),
		TOKEN_LIFE_SPAN: getEnvAsInt("TOKEN_LIFE_SPAN", 1),
	}
}

func getEnvAsString(key, defaultVal string) string {
	if value, isExists := os.LookupEnv(key); isExists {
		return value
	} else {
		return defaultVal
	}
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnvAsString(key, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	} else {
		return defaultVal
	}
}
