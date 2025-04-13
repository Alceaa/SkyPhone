package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alceaa/SkyPhone/models"
)

func GetChats(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)

	var chats []models.Chat
	err = DB.Select(&chats, `SELECT c.id, c.name, c.count
	 FROM chats c 
	 JOIN chats_users cu ON c.id = cu.chatid 
	 WHERE cu.userid = $1 
	 GROUP BY c.id, c.name`, dbUser.ID)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chats)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	var dbUser models.User
	err = DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)

	var contacts []models.Contact
	err = DB.Select(&contacts, `SELECT userid, contactid 
	 FROM contacts 
	 WHERE userid = $1`, dbUser.ID)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
