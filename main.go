package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jarmasp/go-gorm-restapi/db"
	"github.com/jarmasp/go-gorm-restapi/routes"
)

func main() {
	db.DBconnection()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":3000", r)
}
