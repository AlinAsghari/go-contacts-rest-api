package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-contacts/app"
	"go-contacts/controllers"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
)

func main() {

	logger()
	router := mux.NewRouter()

	router.HandleFunc("/api/health", controllers.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthenticationMiddleware) //attach JWT auth middleware
	router.HandleFunc("/jwtapi/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/jwtapi/{id}/contacts/", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(app.BasicAuthMiddleware) //attach JWT auth middleware
	router.HandleFunc("/basicauthapi/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/basicauthapi/{id}/contacts/", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
func logger() {
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}
