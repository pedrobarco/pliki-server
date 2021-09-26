package app

import (
	"net/http"
	"pliki-server/configs"
	"strconv"
)

func CreateServer(cfg *configs.Config) error {
	router := createRouter(cfg)
	port := ":" + strconv.Itoa(cfg.Port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
