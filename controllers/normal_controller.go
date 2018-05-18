package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Prueba2 struct {
	User string
}

var staticFilesPath, _ = filepath.Abs("./src/drivebox/static")

func init() {
	log.SetPrefix("LOG NORMAL CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		ErrorHandler(response, request, http.StatusNotFound)
	} else {
		if CheckAuth(request) {
			http.Redirect(response, request, "/index", http.StatusFound)
		} else {
			http.ServeFile(response, request, "/home/felix/go/src/drivebox/static/auth.html")
		}
	}
}

func ErrorHandler(response http.ResponseWriter, request *http.Request, status int) {
	http.ServeFile(response, request, "/home/felix/go/src/drivebox/static/errors/"+strconv.Itoa(status)+".html")
}
