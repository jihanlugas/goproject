package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/jihanlugas/goproject.git/model"
	"log"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit http://localhost:8010/signin")

	var c model.Credentials
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := c.Signin()

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Invalid Email or Password")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	token, err := c.GenerateToken()

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(token)

	RespondWithJSONAddToken(w, http.StatusOK, map[string]string{"result": "success"}, token)
}