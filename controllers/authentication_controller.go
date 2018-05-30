package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

// LoginHandler ...
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"

	if ComprobarCredenciales(email, pass) {
		SetNewCookie("session", email, response)
		redirectTarget = "/index"
	}

	http.Redirect(response, request, redirectTarget, 302)
}

// LogoutHandler ...
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	SetNewCookie("session", "", response)
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
	if cookie, err := request.Cookie("session"); err == nil && cookie.Value != "" {
		return true
	}
	return false
}

// Homepage ...
func Homepage(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(projectPath + "/views/index.html"))

	email, _ := request.Cookie("session")
	datos := ListarArchivos(email.Value)
	tmpl.Execute(response, datos)
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

	email, _ := request.Cookie("session")

	GuardarArchivo(header.Filename, email.Value)

	http.Redirect(response, request, "/", http.StatusFound)
}

// SetNewCookie ...
func SetNewCookie(cookieName, cookieValue string, response http.ResponseWriter) {
	var cookie *http.Cookie
	if cookieValue != "" {
		cookie = &http.Cookie{
			Name:  cookieName,
			Value: cookieValue,
			Path:  "/",
		}
	} else {
		cookie = &http.Cookie{
			Name:   cookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
	}

	http.SetCookie(response, cookie)
}

// GetCookie ...
func GetCookie(cookieName string, request *http.Request) string {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		log.Println(err)
	}
	return cookie.Value
}

// DownloadHandler ...
func DownloadHandler(response http.ResponseWriter, request *http.Request) {
	email := GetCookie("session", request)
	param := request.URL.Query().Get("name")

	DescifrarArchivo(param, email)

	response.Header().Set("Content-Disposition", "attachment; filename=\""+param+"\"")
	response.Header().Set("Content-Type", request.Header.Get("Content-Type"))

	data, _ := ioutil.ReadFile("./files/" + param)
	http.ServeContent(response, request, param, time.Now(), bytes.NewReader(data))

	deleteFile("files/" + param)
}
