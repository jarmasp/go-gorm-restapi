package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jarmasp/go-gorm-restapi/db"
	"github.com/jarmasp/go-gorm-restapi/models"
	"github.com/jarmasp/go-gorm-restapi/routes"
)

func main() {
	db.DBconnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	// user routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	// task routes
	r.HandleFunc("/users", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
