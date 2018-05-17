package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetNormalRoutes() {
	http.HandleFunc("/404", controllers.NotFound404)
	http.HandleFunc("/401", controllers.Unauthorized401)
}
