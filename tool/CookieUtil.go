package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const COOKIENAME = "gogodancing"

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(COOKIENAME)
	if err == nil {
		context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		return cookie, nil
	} else {
		return nil, err
	}
}
