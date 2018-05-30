package main

import (
	"drivebox/controllers"
	"drivebox/routers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	/* file, e := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Println("Error al abrir el fichero.")
	}
	log.SetOutput(file) */
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {
	controllers.DescifrarArchivo("BBDD.db", "") //descifra la base de datos
	err := os.Chmod("database/BBDD.db", 0777)   //le damos permisos a la bd
	if err != nil {
		fmt.Println(err)
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	mux = routers.InitRoutes(mux)

	go func() {
		if err := http.ListenAndServeTLS(":8080", "./keys/cert.pem", "./keys/key.pem", mux); err != nil {
			fmt.Println(err)
		}
	}()

	<-stopChan              // espera señal SIGINT
	controllers.GestionDB() //al cerrarse el servidor solo quedará la bd cifrada

	/////////////////////////// PRUEBAS //////////////////////////////////
	//controllers.InsertarUsuario("fff", "f")
	//controllers.ListarUsuarios()
}
