package controllers

import (
	"log"
	"net/http"
	"path/filepath"
)

// Prueba ...
type Prueba struct {
	User string
}

func init() {
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

var layoutsPath, _ = filepath.Abs("./src/drivebox/views")

// LoginHandler ...
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"

	if email != "" && pass != "" {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "login",
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		redirectTarget = "/index"
	}

	http.Redirect(response, request, redirectTarget, 302)
}

// LogoutHandler ...
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)

	http.Redirect(response, request, "/", http.StatusFound)
}

// Authentication ...
func Authentication(f http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if CheckAuth(request) {
			f(response, request)
		} else {
			ErrorHandler(response, request, http.StatusUnauthorized)
		}
	}
}

// CheckAuth ...
func CheckAuth(request *http.Request) bool {
	if cookie, err := request.Cookie("session"); err == nil && cookie.Value == "login" {
		return true
	}

	return false
}

// Homepage ...
func Homepage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "/home/felix/go/src/drivebox/views/index.html")
}
