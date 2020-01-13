package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-contacts/models"
	utility "go-contacts/utils"
	"log"
	"net/http"
	"strconv"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //getting value from interface{} type
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utility.Respond(w, utility.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	utility.Respond(w, resp)
}

//var BasicauthCreateContact = func(w http.ResponseWriter, r *http.Request) {
//
//	user := r.Context().Value("user").(uint) //getting value from interface{} type
//	contact := &models.Contact{}
//
//	err := json.NewDecoder(r.Body).Decode(contact)
//	if err != nil {
//		utility.Respond(w, utility.Message(false, "Error while decoding request body"))
//		return
//	}
//
//	contact.UserId = user
//	resp := contact.Create()
//	utility.Respond(w, resp)
//}
var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	//user_id := r.Context().Value("user").(uint) // getting value from interface{} type
	log.Println(" GetContactsFor ....")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	data := models.GetContacts(uint(id))
	resp := utility.Message(true, "success")
	resp["data"] = data
	utility.Respond(w, resp)
}
