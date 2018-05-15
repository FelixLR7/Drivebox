package main

import (
	"net/http"
	/* "fmt" */
	"log"
  /* "github.com/gorilla/mux" */
  "drivebox/routers"
)

func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {
	router := routers.InitRoutes()
	
	http.ListenAndServe(":8080", router)
}