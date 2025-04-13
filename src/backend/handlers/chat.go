package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alceaa/SkyPhone/models"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	var chat models.ChatResponse
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var chatID int
	err := DB.QueryRow("INSERT INTO chats (name, count) VALUES ($1, $2) RETURNING id", chat.Name, len(chat.Users)+1).Scan(&chatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, userID := range chat.Users {
		_, err := DB.Exec("INSERT INTO chats_users (chatid, userid) VALUES ($1, $2)", chatID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	var createdByUser models.User
	err = DB.Get(&createdByUser, `SELECT id, username, name, surname FROM users WHERE username = $1`, chat.CreatedBy)
	_, err = DB.Exec("INSERT INTO chats_users (chatid, userid) VALUES ($1, $2)", chatID, createdByUser.ID)

	w.WriteHeader(http.StatusCreated)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	err := json.NewDecoder(r.Body).Decode(&chat)

	var dbChat models.Chat
	err = DB.Get(&dbChat, `SELECT id, name, count FROM chats WHERE id = $1`, chat.ChatID)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbChat)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	err := json.NewDecoder(r.Body).Decode(&chat)

	var messages []models.MessageResponse

	query := `SELECT m.id, u.username, u.name, u.surname, m.content, m.timestamp::timestamp(0)
	FROM messages m 
	JOIN users u ON m.senderid = u.id 
	WHERE m.chatid = $1 ORDER BY m.timestamp ASC`

	err = DB.Select(&messages, query, chat.ChatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query :=
		`INSERT INTO messages (chatid, senderid, content, timestamp) 
	 	VALUES ($1, $2, $3, NOW())`

	_, err := DB.Exec(query, msg.ChatID, msg.SenderID, msg.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
