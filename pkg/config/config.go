package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() error {
	mainPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("couldn't get mainPath: %v", err)
		return nil
	}

	err = godotenv.Load(mainPath + "/.env")
	if err != nil {
		fmt.Printf("failed godotenv.Load: %v", err)
		return nil
	}

	return nil
}
