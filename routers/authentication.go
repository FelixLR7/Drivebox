package routers
import (
  "github.com/gorilla/mux"
  "drivebox/controllers"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
  router.HandleFunc("/", controllers.AuthHandler)
  router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
  router.HandleFunc("/index", controllers.IndexHandler)
  router.HandleFunc("/css/{file}", controllers.CssHanlder)

  return router
}