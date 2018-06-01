package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
	"time"
)

var doubleAuth = make(map[string]string)

func init() {
	log.SetPrefix("LOG AUTHENTICATION CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

// LoginHandler ...
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")

	if ComprobarCredenciales(email, pass) {
		SetNewCookie("session", email, response)

		token := generateKEY(32)
		doubleAuth[email] = token

		body := "<a href=\"https://localhost:8080/validation?token=" + string(token) + "\">Click aquí</a>"
		SendMail(email, email, "Doble autenticación", body)
		OwnErrorsHandler(response, request, "email")
	}

	OwnErrorsHandler(response, request, "user")
}

// LogoutHandler ...
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	SetNewCookie("session", "", response)
	SetNewCookie("token", "", response)
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
	email, err1 := GetCookie("session", request)
	token, err2 := GetCookie("token", request)

	if err1 == nil && err2 == nil && email != "" && token != "" && doubleAuth[email] == token {
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
func GetCookie(cookieName string, request *http.Request) (string, error) {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return cookie.Value, nil
}

// DownloadHandler ...
func DownloadHandler(response http.ResponseWriter, request *http.Request) {
	email, _ := GetCookie("session", request)
	param := request.URL.Query().Get("name")

	DescifrarArchivo(param, email)

	response.Header().Set("Content-Disposition", "attachment; filename=\""+param+"\"")
	response.Header().Set("Content-Type", request.Header.Get("Content-Type"))

	data, _ := ioutil.ReadFile("./files/" + param)
	http.ServeContent(response, request, param, time.Now(), bytes.NewReader(data))

	deleteFile("files/" + param)
}

// DeleteHandler ...
func DeleteHandler(response http.ResponseWriter, request *http.Request) {
	email, _ := GetCookie("session", request)
	name := request.URL.Query().Get("name")

	EliminarArchivo(name, email)
}

// ValidationHandler ...
func ValidationHandler(response http.ResponseWriter, request *http.Request) {
	email, errEmail := GetCookie("session", request)
	token := request.URL.Query().Get("token")

	if errEmail == nil && doubleAuth[email] == token {
		SetNewCookie("token", token, response)

		http.Redirect(response, request, "/index", http.StatusFound)
	} else {
		ErrorHandler(response, request, http.StatusUnauthorized)
	}
}

// SendMail ...
func SendMail(toNameP, toEmailP, subjectP, bodyP string) {
	fromName := "Drivebox"
	fromEmail := "admin@drivebox.com"
	toNames := []string{toNameP}
	toEmails := []string{toEmailP}
	subject := subjectP
	body := bodyP
	// Build RFC-2822 email
	toAddresses := []string{}
	for i := range toEmails {
		to := mail.Address{toNames[i], toEmails[i]}
		toAddresses = append(toAddresses, to.String())
	}
	toHeader := strings.Join(toAddresses, ", ")
	from := mail.Address{fromName, fromEmail}
	fromHeader := from.String()
	subjectHeader := subject
	header := make(map[string]string)
	header["To"] = toHeader
	header["From"] = fromHeader
	header["Subject"] = subjectHeader
	header["Content-Type"] = `text/html; charset="UTF-8"`
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body
	bMsg := []byte(msg)
	// Send using local postfix service
	c, err := smtp.Dial("localhost:25")
	if err != nil {
		return
	}
	defer c.Close()
	if err = c.Mail(fromHeader); err != nil {
		return
	}
	for _, addr := range toEmails {
		if err = c.Rcpt(addr); err != nil {
			return
		}
	}
	w, err := c.Data()
	if err != nil {
		return
	}
	_, err = w.Write(bMsg)
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	err = c.Quit()
	// Or alternatively, send with remote service like Amazon SES
	// err = smtp.SendMail(addr, auth, fromEmail, toEmails, bMsg)
	// Handle response from local postfix or remote service
	if err != nil {
		return
	}
}
