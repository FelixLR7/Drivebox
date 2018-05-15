package controllers

import (
	"net/http"
	"html/template"
	"log"
	_ "fmt"
	_ "database/sql"
	"path/filepath"
)

type Prueba2 struct {
	User string
}

var staticFilesPath, _ = filepath.Abs("./src/drivebox/static")

func init() {
	log.SetPrefix("LOG NORMAL CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func AuthHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(staticFilesPath + "/auth.html"))
	prueba := Prueba2{User: "a"}
	/* db, err := sql.Open("mysql", "root:123456@/test")

	if err != nil {
		fmt.Println("MAL OPEN")
	}
	defer db.Close() */

	tmpl.Execute(response, prueba)
}

func NotFound404(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(staticFilesPath + "/404.html"))
	tmpl.Execute(response, nil)
}

func Unauthorized401(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(staticFilesPath + "/401.html"))
	tmpl.Execute(response, nil)
}