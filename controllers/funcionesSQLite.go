package controllers

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	DBNAME = "sqlite3"
	DBHOST = "database/BBDD.db"
)

// generateKEY
func generateKEY(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// getKEY
func getKEY(email string) string {
	var key string

	db, _ := sql.Open(DBNAME, DBHOST)
	rows, err := db.Query("SELECT key FROM users WHERE email = '" + email + "';")
	checkErr(err)

	rows.Next()
	rows.Scan(&key)
	rows.Close()

	return key
}

// createDirIfNotExist ...
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		checkErr(err)
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
	db, _ := sql.Open(DBNAME, DBHOST)
	rows, _ := db.Query("SELECT password FROM users WHERE email = '" + email + "';")

	rows.Next()
	rows.Scan(&password)
	rows.Close()

	return password
}

// InsertarUsuario ...
func InsertarUsuario(email, pass string) {
	key := generateKEY(16)
	hash, _ := HashPassword(pass)

	db, _ := sql.Open(DBNAME, DBHOST)
	stmt, err := db.Prepare("INSERT INTO users (email, password, key) values(?,?,?);")
	checkErr(err)

	stmt.Exec(email, hash, key)
	checkErr(err)

	createDirIfNotExist("files/" + email)
}

// EliminarUsuario ...
func EliminarUsuario(email string) {
	db, _ := sql.Open(DBNAME, DBHOST)
	stmt, err := db.Prepare("DELETE FROM users WHERE email = ?;")
	checkErr(err)

	stmt.Exec(email)
	checkErr(err)
}

// ListarUsuarios ...
func ListarUsuarios() {
	db, _ := sql.Open(DBNAME, DBHOST)

	rows, _ := db.Query("SELECT email, password FROM users;")
	var email string
	var password string
	for rows.Next() {
		rows.Scan(&email, &password)
		fmt.Println(email + " - " + password)
	}
	rows.Close()
}

// InsertarArchivos ...
func insertarArchivo(nombre, email string) {
	url := "files/" + email + "/" + nombre

	db, _ := sql.Open(DBNAME, DBHOST)
	stmt, err := db.Prepare("INSERT INTO archivos (nombre, url, emailuser) VALUES(?,?,?);")
	checkErr(err)

	stmt.Exec(nombre, url, email)
	checkErr(err)
}

// EliminarArchivo ...
func eliminarArchivo(archivo, email string) {
	db, _ := sql.Open(DBNAME, DBHOST)
	stmt, err := db.Prepare("DELETE FROM archivos WHERE nombre = ? and emailuser = ?;")
	checkErr(err)

	stmt.Exec(archivo, email)
	checkErr(err)
}

// ListarArchivos ...
func ListarArchivos(useremail string) []string {
	var archivos []string
	var nombre string
	var url string
	var email string

	db, _ := sql.Open(DBNAME, DBHOST)
	rows, _ := db.Query("SELECT * FROM archivos WHERE emailuser = '" + useremail + "';")

	for rows.Next() {
		rows.Scan(&nombre, &url, &email)
		//fmt.Println(nombre + " - " + url + " - " + email)
		archivos = append(archivos, nombre)
	}
	rows.Close()

	return archivos
}

// ComprobarCredenciales
func ComprobarCredenciales(email, pass string) bool {
	hash := datosUsuario(email)
	if hash != "" {
		if CheckPasswordHash(pass, hash) {
			return true
		}
	}
	return false
}
