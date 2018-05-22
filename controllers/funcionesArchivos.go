package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

const (
	KEY = "testtesttesttest"
)

// delete file
func deleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		panic(err)
	}
}
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
func cifrarArchivo(file, key, email string) {
	content, err := readFromFile(file)
	checkErr(err)

	encrypted := encrypt(string(content), key)
	writeToFile(encrypted, "files/"+email+"/"+file+".enc")
}

// Descifrar archivo ...
func descifrarArchivo(file, key, email string) {
	content, err := readFromFile("files/" + email + "/" + file + ".enc")
	checkErr(err)

	decrypted := decrypt(string(content), key)
	writeToFile(decrypted, "files/"+email+"/"+file)
}

// GuardarArchivo  ...
func GuardarArchivo(file, email string) {
	cifrarArchivo(file, KEY, email)
	insertarArchivo(file, email)
}

// DescargarArchivo ...
func DescargarArchivo(file, email string) {
	descifrarArchivo(file, KEY, email)
	//downloadFile(file, "files/"+email)

	deleteFile("files/" + email + "/" + file)
}

// EliminarArchivo ...
func EliminarArchivo(file, email string) {
	eliminarArchivo(file, email)
	deleteFile("files/" + email + "/" + file + ".enc")
}
