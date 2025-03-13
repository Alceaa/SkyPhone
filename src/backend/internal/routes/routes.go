package main

import (
    "github.com/gorilla/mux"
	"skyphone/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")
    router.HandleFunc("/api/login", handlers.Login).Methods("POST")

    router.HandleFunc("/api/contacts", handlers.GetContacts).Methods("GET")
    router.HandleFunc("/api/contacts", handlers.AddContact).Methods("POST")
    router.HandleFunc("/api/contacts/{id}", handlers.DeleteContact).Methods("DELETE")

    router.HandleFunc("/api/chats", handlers.CreateChat).Methods("POST")
    router.HandleFunc("/api/chats/{id}/messages", handlers.GetMessages).Methods("GET")
    router.HandleFunc("/api/chats/{id}/messages", handlers.SendMessage).Methods("POST")
	return router
}
