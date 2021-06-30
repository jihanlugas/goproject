package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jihanlugas/goproject.git/model"
	"log"

	"net/http"
	"strconv"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit http://localhost:8010/projects")
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	projects, err := model.GetProjects(start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusOK, projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit http://localhost:8010/project")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}


	p := model.Project{ID: id}

	if p.ID != 0 {
		if err := p.GetProject(); err != nil {
			switch err {
			case sql.ErrNoRows:
				RespondWithError(w, http.StatusNotFound, "Project not found")
			default:
				RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
	}

	RespondWithSuccess(w, http.StatusOK, p)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit http://localhost:8010/project")
	var p model.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := p.CreateProject(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusCreated, p)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var p model.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	p.ID = id

	if err := p.UpdateProject(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusOK, p)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	p := model.Project{ID: id}
	if err := p.DeleteProject(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusOK, map[string]string{"message": "Data Deleted"})
}


