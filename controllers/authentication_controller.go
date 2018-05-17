package controllers

import (
	"drivebox/services"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	/* "fmt" */ /* "database/sql" */)

type Prueba struct {
	User string
}

func init() {
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

var layoutsPath, _ = filepath.Abs("./src/drivebox/views")

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsPath + "/index.html"))
	email, _ := services.GetCookie("email", request)
	user := Prueba{User: email}
	tmpl.Execute(response, user)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"

	if email != "" && pass != "" {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "",
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		redirectTarget = "/index"
	}

	http.Redirect(response, request, redirectTarget, 302)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)

	http.Redirect(response, request, "/", http.StatusFound)
}

func CheckAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if _, err := request.Cookie("session"); err == nil {
			f(response, request)
		} else {
			ErrorHandler(response, request, http.StatusUnauthorized)
		}
	}
}
