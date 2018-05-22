package routers

import (
	"net/http"
)

func InitRoutes(mux *http.ServeMux) *http.ServeMux {
	mux = SetAuthenticationRoutes(mux)
	mux = SetNormalRoutes(mux)

	return mux
}
