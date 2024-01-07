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
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/user", routes.DeleteUsersHandler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
