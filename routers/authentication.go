package routers

import (
	"drivebox/controllers"
	"net/http"
)

// SetAuthenticationRoutes ...
func SetAuthenticationRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/login", controllers.Authentication(controllers.LoginHandler))
	mux.HandleFunc("/logout", controllers.Authentication(controllers.LogoutHandler))
	mux.HandleFunc("/upload", controllers.Authentication(func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.UploadPageHandler(response, request)
		} else {
			controllers.UploadHandler(response, request)
		}
	}))
	mux.HandleFunc("/download/{file}", controllers.Authentication(controllers.DownloadHandler))
	mux.HandleFunc("/index", controllers.Authentication(controllers.Homepage))

	return mux
}
