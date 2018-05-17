package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func listarUsuarios() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// listar usuarios
	results, err := db.Query("SELECT email, pass FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Email, &user.Pass)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Email)
		fmt.Println(user.Pass)
	}
}
func insertarUsuario(email string, pass string) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO users VALUES('" + email + "','" + pass + "');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Usuario insertado correctamente")
}

func main() {
	//insertarUsuario("a@a.a", "a")
	listarUsuarios()
}
