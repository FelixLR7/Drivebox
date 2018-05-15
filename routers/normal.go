package routers
import (
  "github.com/gorilla/mux"
  "drivebox/controllers"
  "net/http"
)

func SetNormalRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.AuthHandler)
  router.HandleFunc("/css/{file}", controllers.CssHandler)
  
  return router
}