package controllers

import (
	"net/http"
	"html/template"
	"path/filepath"
	"drivebox/services"
	"log"
)

type Prueba struct {
	User string
}

func init(){
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

var layoutsAbsPath, _ = filepath.Abs("./src/drivebox/views")

func AuthHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/auth.html"))
	tmpl.Execute(response, nil)
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/index.html"))
	user := Prueba{User: services.GetUserName(request)}
	tmpl.Execute(response, user)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	
	if email != "" && pass != "" {
		services.NewCookie("email", email, response)
		redirectTarget = "/index"
	}
	
	http.Redirect(response, request, redirectTarget, 302)
}

