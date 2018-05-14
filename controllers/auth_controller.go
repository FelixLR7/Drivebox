package controllers

import (
	"net/http"
	"strings"
	"html/template"
	"path/filepath"
	"fmt"
)

var layoutsAbsPath, _ = filepath.Abs("./src/drivebox/views")

func AuthHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/auth.html"))
	tmpl.Execute(response, nil)
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/index.html"))
	tmpl.Execute(response, nil)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	
	if name != "" && pass != "" {
		/* setSession(name, response) */
		redirectTarget = "/index"
	}
	
	http.Redirect(response, request, redirectTarget, 302)
}

func CssHanlder(response http.ResponseWriter, request *http.Request) {
  path := strings.Split(request.URL.Path, "/")
	if len(path) == 0 {
		fmt.Println("a")
	}
}