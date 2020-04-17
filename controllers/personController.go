package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sudhanshu-chauhan/gopad/db"
	"github.com/sudhanshu-chauhan/gopad/models"
)

func PostPerson(w http.ResponseWriter, r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)
	responseEncoder := json.NewEncoder(w)
	var person models.Person
	err := bodyDecoder.Decode(&person)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	db.CreatePerson(&person)
	responseEncoder.Encode(&person)

}
