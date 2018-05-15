package routers

import (
  "github.com/gorilla/mux"
  "drivebox/controllers"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
  router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
  router.HandleFunc("/index", controllers.CheckAuth(controllers.IndexHandler))
  router.HandleFunc("/", controllers.AuthHandler)
  return router
}