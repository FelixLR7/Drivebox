package main

import (
	"drivebox/routers"
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {
	mux := http.NewServeMux()
	mux = routers.InitRoutes(mux)

	srv := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		if err := srv.ListenAndServeTLS("keys/cert.pem", "keys/key.pem"); err != nil {
			log.Println("Escuchando")
		}
	}()

	/////////////////////////// PRUEBAS //////////////////////////////////

}
