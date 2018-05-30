package main

import (
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
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	mux = routers.InitRoutes(mux)

	go func() {
		if err := http.ListenAndServeTLS(":8080", "./keys/cert.pem", "./keys/key.pem", mux); err != nil {
			fmt.Println(err)
		}
	}()

	<-stopChan // espera señal SIGINT

	/////////////////////////// PRUEBAS //////////////////////////////////

}
