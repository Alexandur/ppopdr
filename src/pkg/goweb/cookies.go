package goweb

/*
	Interim code to make cookies work on appengine
	for release r58.1
*/

import (
	"http"
	"strings"
)

func parseCookieHeader(cookieHeader string, request *http.Request) {
	// break up into individual cookies
	cookieStrings := strings.Split(cookieHeader, ";")

	for _, cookieString := range cookieStrings {
		var cookie *http.Cookie = new(http.Cookie)

		cookieStringSegments := strings.Split(strings.TrimSpace(cookieString), "=")

		cookie.Name = cookieStringSegments[0]
		cookie.Value = cookieStringSegments[1]

		request.AddCookie(cookie)
	}
}




