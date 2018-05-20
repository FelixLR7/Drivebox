package controllers

import (
	"database/sql"
	"fmt"
	"os"
	// a
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "testdb"
	DB_USER = "root"
	DB_PASS = "admin"
)

// User ...
type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// Archivo ...
type Archivo struct {
	Nombre string `json:"nombre"`
	Url    string `json:"url"`
}

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// createDirIfNotExist ...
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ListarUsuarios() {
	db := dbConn()
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
	defer db.Close()
}

// InsertarUsuario ...
func InsertarUsuario(email, pass string) {
	db := dbConn()

	hash, _ := HashPassword(pass)
	insert, err := db.Query("INSERT INTO users VALUES('" + email + "','" + hash + "');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	createDirIfNotExist(email)

	fmt.Println("Usuario insertado correctamente")
}

// ListarArchivos ...
func ListarArchivos(emailUser string) {
	db := dbConn()

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
	defer db.Close()
}

// InsertarArchivo ...
func InsertarArchivo(emailUser, nombre, url string) {
	db := dbConn()

	insert, err := db.Query("INSERT INTO archivo VALUES('" + nombre + "','" + url + "','" + emailUser + "');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Archivo insertado correctamente")
}

// EliminarArchivo ...
func EliminarArchivo(emailUser, nombre string) {
	db := dbConn()

	insert, err := db.Query("DELETE FROM archivo WHERE emailuser='" + emailUser + "' and nombre='" + nombre + "';")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Archivo borrado correctamente")
}
