package controllers

import (
	"encoding/json"
	"go-contacts/models"
	utility "go-contacts/utils"
	"log"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	log.Println(" CreateAccount ....")
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		utility.Respond(w, utility.Message(false, "Invalid request "+err.Error()))
		return
	}

	resp := account.Create() //Create account
	utility.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	log.Println(" Authenticate ....")

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		utility.Respond(w, utility.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	utility.Respond(w, resp)
}
