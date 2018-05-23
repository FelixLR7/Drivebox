package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetAuthenticationRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/login", controllers.LoginHandler)
	mux.HandleFunc("/logout", controllers.LogoutHandler)
	mux.HandleFunc("/upload", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.UploadPageHandler(response, request)
		} else {
			controllers.UploadHandler(response, request)
		}
	})
	mux.HandleFunc("/download/{file}", controllers.DownloadHandler)
	mux.HandleFunc("/index", controllers.Authentication(controllers.Homepage))

	return mux
}
