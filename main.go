package main

import (
	"drivebox/controllers"
	"log"
	/* "github.com/gorilla/mux" */ /* "drivebox/routers" */)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Server started")
}

func main() {
	//routers.InitRoutes()
	//http.ListenAndServe(":8080", nil)

	/////////////////////////// PRUEBAS //////////////////////////////////
	controllers.EliminarArchivo("g.txt", "a@a.a")
	//a := controllers.ListarArchivos("a@a.a")
}
