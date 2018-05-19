package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetAuthenticationRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/logout", controllers.LogoutHandler)
	http.HandleFunc("/upload", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.UploadPageHandler(response, request)
		} else {
			controllers.UploadHandler(response, request)
		}
	})
	http.HandleFunc("/index", controllers.Authentication(controllers.Homepage))
}
