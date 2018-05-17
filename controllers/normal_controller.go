package controllers

import (
	_ "database/sql"
	_ "fmt"
	"log"
	"net/http"
	"path/filepath"
)

type Prueba2 struct {
	User string
}

var staticFilesPath, _ = filepath.Abs("./src/drivebox/static")

func init() {
	log.SetPrefix("LOG NORMAL CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func ErrorHandler(response http.ResponseWriter, request *http.Request, status int) {
	response.WriteHeader(status)
	if status == http.StatusNotFound {
		http.ServeFile(response, request, "/home/felix/go/src/drivebox/static/errors/404.html")
	}
}
