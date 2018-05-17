package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func HashPassword(password string) (string, error) { //cifra el pass
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool { //devuelve true si el hash es igual a la pass
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func listarUsuarios() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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

		fmt.Println(user.Email + ": " + user.Pass)
	}
}
func insertarUsuario(email string, pass string) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	hash, _ := HashPassword(pass)
	insert, err := db.Query("INSERT INTO users VALUES('" + email + "','" + hash + "');")
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
