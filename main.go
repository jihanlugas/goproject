package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jihanlugas/goproject.git/controller"
	"log"
	"net/http"
)

func main()  {
	log.Println("Listening server at http://localhost:8010")
	router := mux.NewRouter()

	router.HandleFunc("/signin", controller.Signin).Methods("POST")

	router.HandleFunc("/projects", controller.GetProjects).Methods("GET")
	router.HandleFunc("/project/{id:[0-9]+}", controller.GetProject).Methods("GET")
	router.HandleFunc("/project", controller.CreateProject).Methods("POST")
	router.HandleFunc("/project/{id:[0-9]+}", controller.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/{id:[0-9]+}", controller.DeleteProject).Methods("DELETE")

	router.HandleFunc("/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id:[0-9]+}", controller.DeleteTask).Methods("DELETE")

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", controller.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8010", router))
}
