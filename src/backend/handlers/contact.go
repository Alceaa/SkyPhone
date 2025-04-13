package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alceaa/SkyPhone/models"
)

func AddContact(w http.ResponseWriter, r *http.Request) {
	var contact models.ContactResponse
	err := json.NewDecoder(r.Body).Decode(&contact)

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", contact.AddBy)

	_, err = DB.Exec("INSERT INTO contacts (userid, contactid) VALUES ($1, $2)", dbUser.ID, contact.ContactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	var contact models.ContactResponse
	err := json.NewDecoder(r.Body).Decode(&contact)

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", contact.AddBy)

	_, err = DB.Exec("DELETE FROM contacts WHERE userid=$1 AND contactID=$2", dbUser.ID, contact.ContactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUserByContactID(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	var user models.User
	err = DB.Get(&user, `SELECT id, username, name, surname FROM users WHERE id = $1`, contact.ContactID)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetContactsUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)

	var users []models.User
	err = DB.Select(&users, `SELECT u.id, u.username, u.name, u.surname
		FROM contacts c
		JOIN users u ON c.contactId = u.id
		WHERE c.userId = $1`, dbUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetNotContactsUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)

	var users []models.User
	err = DB.Select(&users, `SELECT u.id, u.username, u.name, u.surname
		FROM users u
		LEFT JOIN contacts c ON u.id = c.contactId AND c.userId = $1
		WHERE c.contactId IS NULL AND u.id <> $1`, dbUser.ID)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
