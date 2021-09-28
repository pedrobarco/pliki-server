package app

import (
	"net/http"
	"os"
	"pliki-server/configs"
	"strconv"
)

func CreateServer(cfg *configs.Config) error {
	err := os.MkdirAll(cfg.FileServerRoot, 0755)
	if err != nil {
		return err
	}

	router := createRouter(cfg)
	port := ":" + strconv.Itoa(cfg.Port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		return err
	}

	return nil
}
