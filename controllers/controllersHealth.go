package controllers

import (
	utility "go-contacts/utils"
	"log"
	"net/http"
)

var HealthCheckHandler = func(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	log.Println(" HealthCheckHandler ....")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	//io.WriteString(w, `{"alive": true}`)
	resp := utility.Message(true, "OK")
	utility.Respond(w, resp)
}
