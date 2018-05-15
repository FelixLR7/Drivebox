package routers
import (
  "github.com/gorilla/mux"
  "drivebox/controllers"
)

func SetNormalRoutes(router *mux.Router) *mux.Router {
  router.HandleFunc("/404", controllers.NotFound404)
  router.HandleFunc("/401", controllers.Unauthorized401)
  router.HandleFunc("/css/{file}", controllers.CssHandler) 
  
  return router
}