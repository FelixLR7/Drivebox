package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetNormalRoutes() {
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir("/home/felix/go/src/drivebox/static/css/"))))
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			controllers.ErrorHandler(response, request, http.StatusNotFound)
		} else {
			http.ServeFile(response, request, "/home/felix/go/src/drivebox/static/auth.html")
		}
	})
}
