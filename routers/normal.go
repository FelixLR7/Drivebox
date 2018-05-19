package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetNormalRoutes() {
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir("/home/felix/go/src/drivebox/static/css/"))))
	/* http.HandleFunc("/register", controllers.RegisterPageHandler) */
	http.HandleFunc("/register", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.RegisterPageHandler(response, request)
		} else {
			controllers.RegisterHandler(response, request)
		}
	})
	http.HandleFunc("/", controllers.IndexHandler)
}
