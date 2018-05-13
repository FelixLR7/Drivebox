package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	/* "fmt" */
	"log"
  "github.com/gorilla/mux"
  
)

func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

/* func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
} */

func authHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/auth.html"))
	tmpl.Execute(response, nil)
	log.Printf("Se ejecuta")
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/index.html"))
	tmpl.Execute(response, nil)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	
	if name != "" && pass != "" {
		/* setSession(name, response) */
		redirectTarget = "/index"
	}
	
	http.Redirect(response, request, redirectTarget, 302)
}

var router = mux.NewRouter()
var layoutsAbsPath, _ = filepath.Abs("./src/drivebox/layouts")

func main() {
	router.HandleFunc("/", authHandler)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/index", indexHandler)
	router.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no auth required\n"))
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}