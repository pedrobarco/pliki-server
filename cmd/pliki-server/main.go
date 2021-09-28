package main

import (
	"log"
	"os"
	"pliki-server/configs"
	"pliki-server/web/app"
	"strconv"
)

func main() {
	cfg := configs.GetDefaultConfig()

	envPort := os.Getenv("PLIKI_PORT")
	envFileServerRoot := os.Getenv("PLIKI_FILE_SERVER_ROOT")
	envMaxFileSizeMb := os.Getenv("PLIKI_MAX_FILE_SIZE_MB")

	if envPort != "" {
		cfg.Port, _ = strconv.Atoi(envPort)
	}

	if envFileServerRoot != "" {
		cfg.FileServerRoot = envFileServerRoot
	}

	if envMaxFileSizeMb != "" {
		cfg.MaxFileSizeMb, _ = strconv.ParseInt(envMaxFileSizeMb, 10, 64)
	}

	err := app.CreateServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
