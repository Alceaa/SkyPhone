package handlers

import (
    "encoding/json"
    "net/http"
)

func AddContact(w http.ResponseWriter, r *http.Request) {
    // Логика добавления контакта в БД
    
    w.WriteHeader(http.StatusCreated)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
    // Логика удаления контакта из БД
    
    w.WriteHeader(http.StatusNoContent)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
    // Логика получения списка контактов из БД
    
    contacts := []models.Contact{}
    json.NewEncoder(w).Encode(contacts)
}