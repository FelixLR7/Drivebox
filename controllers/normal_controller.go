package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// Prueba2 ...
type Prueba2 struct {
	User string
}

var projectPath, _ = filepath.Abs("./src/drivebox")
var staticFilesPath = projectPath + "/static"

func init() {
	log.SetPrefix("LOG NORMAL CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

// IndexHandler ...
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		ErrorHandler(response, request, http.StatusNotFound)
	} else {
		if CheckAuth(request) {
			http.Redirect(response, request, "/index", http.StatusFound)
		} else {
			http.ServeFile(response, request, staticFilesPath+"/auth.html")
		}
	}
}

// ErrorHandler ...
func ErrorHandler(response http.ResponseWriter, request *http.Request, status int) {
	http.ServeFile(response, request, staticFilesPath+"/errors/"+strconv.Itoa(status)+".html")
}

// RegisterPageHandler ...
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, staticFilesPath+"/register.html")
}

// RegisterHandler ...
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "/", http.StatusFound)
}
