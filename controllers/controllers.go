package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm-crud/db"
	"gorm-crud/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.Users

	json.NewDecoder(r.Body).Decode(&user)
	db.Db.Create(&user)
	json.NewEncoder(w).Encode(&user)

}

func checkIfUserExists(userId string) bool {
	var user models.Users
	db.Db.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("User not found")
		return
	}

	var user models.Users
	db.Db.First(&user, userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	db.Db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User doesn't exist")
		return
	}
	var user models.Users
	db.Db.Delete(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]

	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No user found")
		return
	}

	var user models.Users
	db.Db.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	db.Db.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

}
