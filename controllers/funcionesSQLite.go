package controllers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	DB_NAME = "sqlite3"
	DB_HOST = "database/BBDD.db"
)

// createDirIfNotExist ...
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
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

func datosUsuario(email string) string {
	var password string
	database, _ := sql.Open(DB_NAME, DB_HOST)
	rows, _ := database.Query("SELECT password FROM users WHERE email = '" + email + "';")

	rows.Next()
	rows.Scan(&password)

	return password
}

// Insertar Usuario ...
func InsertarUsuario(email, pass string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("INSERT INTO users (email, password) values(?,?)")
	checkErr(err)

	hash, _ := HashPassword(pass)
	stmt.Exec(email, hash)
	checkErr(err)
	defer db.Close()

	createDirIfNotExist("files/" + email)
}

// Eliminar Usuario ...
func EliminarUsuario(email string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("DELETE FROM users WHERE email = ?")
	checkErr(err)

	stmt.Exec(email)
	checkErr(err)
}

// Listar Usuarios ...
func ListarUsuarios() {
	database, _ := sql.Open(DB_NAME, DB_HOST)

	rows, _ := database.Query("SELECT * FROM users")
	var email string
	var password string
	for rows.Next() {
		rows.Scan(&email, &password)
		fmt.Println(email + " - " + password)
	}
}

// Insertar Archivos ...
func insertarArchivo(nombre, email string) {
	url := "files/" + email + "/" + nombre

	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("INSERT INTO archivos (nombre, url, emailuser) values(?,?,?)")
	checkErr(err)

	stmt.Exec(nombre, url, email)
	checkErr(err)
}

// Eliminar Archivo ...
func eliminarArchivo(archivo, email string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("DELETE FROM archivos WHERE nombre = ? and emailuser = ?")
	checkErr(err)

	stmt.Exec(archivo, email)
	checkErr(err)
}

// Listar Archivos ...
func ListarArchivos(useremail string) []string {
	var archivos []string
	var nombre string
	var url string
	var email string

	database, _ := sql.Open(DB_NAME, DB_HOST)
	rows, _ := database.Query("SELECT * FROM archivos WHERE emailuser = '" + useremail + "';")

	for rows.Next() {
		rows.Scan(&nombre, &url, &email)
		//fmt.Println(nombre + " - " + url + " - " + email)
		archivos = append(archivos, nombre)
	}

	return archivos
}

// ComprobarCredenciales
func ComprobarCredenciales(email, pass string) bool {
	hash := datosUsuario(email)
	if hash != "" {
		if CheckPasswordHash(pass, hash) {
			return true
		} else {
			fmt.Println("Contraseña incorrecta !!!")
			return false
		}
	} else {
		fmt.Println("El usuario no existe !!!")
		return false
	}
}
