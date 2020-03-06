package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateDocument = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	document := &models.Document{}

	err := json.NewDecoder(r.Body).Decode(document)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	document.UserId = user
	resp := document.Create()
	u.Respond(w, resp)
}

var UpdateDocument = func(w http.ResponseWriter, r *http.Request) {

	document := &models.Document{}
	
	err := json.NewDecoder(r.Body).Decode(document)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.UpdateDocument(document.ID, document)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteDocument = func(w http.ResponseWriter, r *http.Request) {

	document := &models.Document{}
	
	err := json.NewDecoder(r.Body).Decode(document)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.DeleteDocument(document.ID, document)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetDocumentsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetDocuments(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
