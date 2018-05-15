package services

import(
	"github.com/gorilla/securecookie"
	"net/http"
	"log"
	"errors"
)
	
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Crea una cookie dado un nombre y un valor
func NewCookie(cookieName string, cookieValue string, response http.ResponseWriter) {
	if encoded, err := cookieHandler.Encode(cookieName, cookieValue); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		log.Println("Nueva cookie: " + cookieName)
	}
}

// Borra una cookie dado su nombre
func ClearSession(cookieName string, response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// Devuelve el valor de una cookie dado su nombre
func GetCookie(cookieName string, request *http.Request) (cookieValue string, err error) {
	if cookie, err := request.Cookie(cookieName); err == nil {		
		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue); err == nil {
			return cookieValue, nil
		} else {
			return "", errors.New("No estás autorizado")
		}		
	} else {
		return "", errors.New("No estás autorizado")
	}
}