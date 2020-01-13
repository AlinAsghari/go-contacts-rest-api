package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)
var BasicAuthMiddleware = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path                               //current request path
		if !strings.HasPrefix(requestPath, "/basicauthapi"){
			next.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		userId , _  := getUserIdByUserName(user)
		ctx := context.WithValue(r.Context(), "user", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	});
}
func getUserIdByUserName(username string ) ( uint , error ) {
	if username == "root" {
		return 1 , nil //
	}
	return 1 , nil
}
func checkUsernameAndPassword(username, password string) bool {
	return username == "root" && password == "linux2013"
}

