package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetNormalRoutes() {
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir("/home/felix/go/src/drivebox/static/css/"))))
	http.HandleFunc("/", controllers.IndexHandler)
}
