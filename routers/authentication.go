package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetAuthenticationRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/index", controllers.CheckAuth(controllers.IndexHandler))
}
