package app

import (
	utility "go-contacts/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		utility.Respond(w, utility.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
