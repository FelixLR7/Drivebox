package main

import (
	"drivebox/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.IndexHandler)

	/* srv := &http.Server{Addr: ":8080", Handler: mux} */

	go func() {
		/* if err := srv.ListenAndServeTLS("./keys/cert.pem", "./keys/key.pem"); err != nil {
			fmt.Println(err)
		} */
		if err := http.ListenAndServe(":8080", mux); err != nil {
			fmt.Println(err)
		}
	}()

	<-stopChan // espera señal SIGINT

	/////////////////////////// PRUEBAS //////////////////////////////////

}
