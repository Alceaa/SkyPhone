package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Alceaa/SkyPhone/db"
	"github.com/Alceaa/SkyPhone/models"
	"github.com/Alceaa/SkyPhone/utils"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB = db.DB

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)
	if dbUser.Username != "" {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	_, err = DB.NamedExec(`INSERT INTO users (username, password, name, surname) VALUES (:username, :password, :name, :surname)`, &user)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful!"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)
	if dbUser.Username != user.Username {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful!"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)
	if dbUser.Username != user.Username {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dbUser)
}
