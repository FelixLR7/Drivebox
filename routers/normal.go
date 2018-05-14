package routers
import (
  "github.com/gorilla/mux"
  "drivebox/controllers"
  "net/http"
)

func SetNormalRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.AuthHandler)
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  
  return router
}