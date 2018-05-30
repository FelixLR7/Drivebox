package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

// delete file
func deleteFile(path string) {
	var err = os.Remove(path)
	checkErr(err)
}

// checkErr
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Decrypt
func decrypt(cipherstring string, keystring string) string {
	ciphertext := []byte(cipherstring)
	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	checkErr(err)

	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

// Encrypt
func encrypt(plainstring, keystring string) string {
	plaintext := []byte(plainstring)
	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	checkErr(err)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

// WriteToFile
func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

// ReadFromFile
func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

// Cifrar archivo ...
func CifrarArchivo(file, key, email string) {
	var readRuta string
	var writeRuta string

	if file == "BBDD.db" { //para cifrar la base de datos
		readRuta = "database/" + file
		writeRuta = readRuta
		key = "testtesttesttest"
	} else {
		writeRuta = "files/" + email + "/" + file
		readRuta = "files/" + file
	}
	content, err := readFromFile(readRuta)
	checkErr(err)

	encrypted := encrypt(string(content), key)
	writeToFile(encrypted, writeRuta+".enc")

}

// Descifrar archivo ...
func DescifrarArchivo(file, email string) {
	var key string
	var readRuta string
	var writeRuta string

	if file == "BBDD.db" {
		key = "testtesttesttest"
		readRuta = "database/" + file
		writeRuta = readRuta
	} else {
		key = getKEY(email)
		readRuta = "files/" + email + "/" + file
		writeRuta = "files/" + file
	}

	content, err := readFromFile(readRuta + ".enc")
	checkErr(err)

	decrypted := decrypt(string(content), key)
	writeToFile(decrypted, writeRuta)
}

// GuardarArchivo  ...
func GuardarArchivo(file, email string) {
	insertarArchivo(file, email)
	key := getKEY(email)
	CifrarArchivo(file, key, email)
	deleteFile("files/" + file)

}

// EliminarArchivo ...
func EliminarArchivo(file, email string) {
	eliminarArchivo(file, email)
	deleteFile("files/" + email + "/" + file + ".enc")
}

//GestionDB
func GestionDB() {
	deleteFile("database/BBDD.db.enc") //borramos la bd cifrada
	CifrarArchivo("BBDD.db", "", "")   //ciframos la nueva bd
	deleteFile("database/BBDD.db")     //borramos la bd sin cifrar
}
