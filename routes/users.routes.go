package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jarmasp/go-gorm-restapi/db"
	"github.com/jarmasp/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	CreatedUser := db.DB.Create(&user)
	err := CreatedUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
