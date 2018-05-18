package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
type Archivo struct {
	Nombre string `json:"nombre"`
	Url    string `json:"url"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
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
func insertarUsuario(email, pass string) {
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
	createDirIfNotExist(email)

	fmt.Println("Usuario insertado correctamente")
}
func listarArchivos(emailUser string) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT nombre, url FROM archivo WHERE emailuser='" + emailUser + "';")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var archivo Archivo

		err = results.Scan(&archivo.Nombre, &archivo.Url)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("NOMBRE: " + archivo.Nombre + " - URL: " + archivo.Url)
	}
}
func insertarArchivo(emailUser, nombre, url string) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO archivo VALUES('" + nombre + "','" + url + "','" + emailUser + "');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Archivo insertado correctamente")
}
func eliminarArchivo(emailUser, nombre string) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query("DELETE FROM archivo WHERE emailuser='" + emailUser + "' and nombre='" + nombre + "';")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Archivo borrado correctamente")
}

func main() {
	insertarUsuario("b@b.b", "hola")
	listarUsuarios()

}
