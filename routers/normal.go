package routers

import (
	"drivebox/controllers"
	"net/http"
)

// SetNormalRoutes ...
func SetNormalRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/register", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.RegisterPageHandler(response, request)
		} else {
			controllers.RegisterHandler(response, request)
		}
	})
	mux.HandleFunc("/", controllers.IndexHandler)

	return mux
}
