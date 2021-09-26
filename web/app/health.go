package app

import (
	"net/http"
	"pliki-server/configs"

	"github.com/julienschmidt/httprouter"
)

func healthHandler(cfg *configs.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}
}
