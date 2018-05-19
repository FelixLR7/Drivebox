package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Prueba ...
type Prueba struct {
	User string
}

func init() {
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

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
	http.ServeFile(response, request, projectPath+"/views/index.html")
}

// UploadPageHandler ...
func UploadPageHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, projectPath+"/views/upload.html")
}

// UploadHandler ...
func UploadHandler(response http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Destination
	dst, err := os.Create(projectPath + "/files/" + header.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(response, request, "/", http.StatusFound)
}
