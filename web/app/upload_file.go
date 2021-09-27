package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"pliki-server/configs"
	"time"

	"github.com/julienschmidt/httprouter"
)

func uploadFileHandler(cfg *configs.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := r.ParseMultipartForm(cfg.MaxFileSizeMb << 20)
		if err != nil {
			log.Printf("Error calculating file size: %s", err)
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			// TODO: handle error with 500
			log.Printf("Error parsing file: %s", err)
			return
		}
		defer file.Close()

		ts := time.Now()
		fileName := fmt.Sprintf("%d%02d%02d%02d%02d_%s",
			ts.Year(), ts.Month(), ts.Day(),
			ts.Hour(), ts.Minute(),
			handler.Filename)
		filePath := cfg.FileServerRoot + fileName

		dst, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			// TODO: handle error with 500
			log.Printf("Error opening file: %s", err)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			// TODO: handle error with 500
			log.Printf("Error saving file: %s", err)
			return
		}

		// TODO: check if http or https
		furl := fmt.Sprintf("http://%s/files/%s", r.Host, fileName)

		fmt.Fprint(w, furl)
	}
}
