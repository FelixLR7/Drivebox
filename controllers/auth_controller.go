package controllers

import (
	"net/http"
	"html/template"
	"path/filepath"
	"drivebox/services"
	"log"
	"fmt"
)

type Prueba struct {
	User string
}

func init() {
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
	email, _ := services.GetCookie("email", request)
	user := Prueba{User: email}
	tmpl.Execute(response, user)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	
	if email != "" && pass != "" {
		services.NewCookie("email", email, response)
		services.NewCookie("password", pass, response)
		redirectTarget = "/index"
	}
	
	http.Redirect(response, request, redirectTarget, 302)
}

func CheckAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		_, errEmail := services.GetCookie("email", request)
		_, errPassword := services.GetCookie("password", request)

		if errEmail == nil && errPassword == nil {
			f(response, request)
		} else {
			fmt.Println("NO ESTÁS AUTORIZADO")
		}
	}
}

