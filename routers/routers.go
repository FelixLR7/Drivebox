package routers

import (
	"net/http"
)

// InitRoutes ...
func InitRoutes(mux *http.ServeMux) *http.ServeMux {
	mux = SetAuthenticationRoutes(mux)
	mux = SetNormalRoutes(mux)

	return mux
}
