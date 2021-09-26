package main

import (
	"log"
	"pliki-server/configs"
	"pliki-server/web/app"
)

func main() {
	cfg := configs.GetDefaultConfig()
	err := app.CreateServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
