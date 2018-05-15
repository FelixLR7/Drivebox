package controllers

import (
	"net/http"
	"html/template"
	"log"
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Prueba2 struct {
	User string
}

func init() {
	log.SetPrefix("LOG NORMAL CONTROLLER: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func AuthHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutsAbsPath + "/auth.html"))
	prueba := Prueba2{User: "a"}
	db, err := sql.Open("mysql", "root:123456@/test")

	if err != nil {
		fmt.Println("MAL OPEN")
	}
	defer db.Close()

	tmpl.Execute(response, prueba)
}