package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jihanlugas/goproject.git/model"
	"log"
	"net/http"
	"strconv"
)

func CreateTask(w http.ResponseWriter, r * http.Request){
	log.Println("Hit http://localhost:8010/task")
	var t model.ProjectTask
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	defer r.Body.Close()

	if err := t.CreateTask(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusCreated, t)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	t := model.ProjectTask{ID: id}
	if err := t.DeleteTask(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusOK, map[string]string{"message": "Data Deleted"})
}