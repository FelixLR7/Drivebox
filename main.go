package main

import (
	"drivebox/routers"
	"log"
	"net/http"
	/* "github.com/gorilla/mux" */ /* "drivebox/routers" */)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {

	routers.InitRoutes()
	http.ListenAndServe(":8080", nil)

	/////////////////////////// PRUEBAS //////////////////////////////////
	//file := "b.pdf.enc"
	//key := "testtesttesttest"

	/*	CIFRAR
		content, err := readFromFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		encrypted := encrypt(string(content), key)
		writeToFile(encrypted, file+".enc")
	*/
	/*	DESCIFRAR
		content, err := readFromFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		decrypted := decrypt(string(content), key)
		writeToFile(decrypted, file[:len(file)-4])
	*/
}
