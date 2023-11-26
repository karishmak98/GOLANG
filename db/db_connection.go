package db

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "PRACTICE"

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `./config/.env`)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error loading .env file")
	}
}
