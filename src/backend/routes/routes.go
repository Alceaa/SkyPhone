package routes

import (
	"github.com/Alceaa/SkyPhone/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/user", handlers.GetUser).Methods("POST")
	r.HandleFunc("/api/users", handlers.GetUsers)
	r.HandleFunc("/api/users/contacts", handlers.GetContactsUsers).Methods("POST")
	r.HandleFunc("/api/users/notcontacts", handlers.GetNotContactsUsers).Methods("POST")
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/chats", handlers.GetChats).Methods("POST")
	r.HandleFunc("/api/chat", handlers.GetChat).Methods("POST")
	r.HandleFunc("/api/create/chat", handlers.CreateChat).Methods("POST")
	r.HandleFunc("/api/contacts", handlers.GetContacts).Methods("POST")
	r.HandleFunc("/api/contact", handlers.GetUserByContactID).Methods("POST")
	r.HandleFunc("/api/contact", handlers.DeleteContact).Methods("DELETE")
	r.HandleFunc("/api/add/contact", handlers.AddContact).Methods("POST")
	r.HandleFunc("/api/messages", handlers.GetMessages).Methods("POST")
	r.HandleFunc("/api/send", handlers.SendMessage).Methods("POST")
	return r
}
