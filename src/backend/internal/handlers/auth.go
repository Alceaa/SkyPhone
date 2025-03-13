package handlers

import (
    "encoding/json"
    "net/http"
    "skyphone/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil || user.Username == "" || user.Password == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    
    // Логика сохранения пользователя в БД
    
    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    
    // Логика проверки пользователя и выдачи JWT токена
    
    w.WriteHeader(http.StatusOK)
}