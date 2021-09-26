package app

import (
	"net/http"
	"pliki-server/configs"

	"github.com/julienschmidt/httprouter"
)

func createRouter(cfg *configs.Config) *httprouter.Router {
	router := httprouter.New()

	// register handler for health status
	router.GET("/health", healthHandler(cfg))

	// register handler for uploading files
	router.POST("/", uploadFileHandler(cfg))

	// register handler for file server
	router.ServeFiles("/files/*filepath", http.Dir(cfg.FileServerRoot))

	return router
}
