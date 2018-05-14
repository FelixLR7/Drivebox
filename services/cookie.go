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


func NewCookie(cookieName string, cookieValue string, response http.ResponseWriter) {
	log.Println("a")
	if encoded, err := cookieHandler.Encode(cookieName, cookieValue); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(cookieName string, response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetCookie(cookieName string, request *http.Request) (cookieValue string, err error) {
	if cookie, err := request.Cookie(cookieName); err == nil {
		cookieAux := ""
		
		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieAux); err == nil {
			cookieValue = cookieAux
		} else {
			return "", errors.New("No estás autorizado")
		}
	}
	return cookieValue, nil
}