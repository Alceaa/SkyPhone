package models

type Message struct {
    ID        string json:"id"
    ChatID    string json:"chat_id"
    SenderID  string json:"sender_id"
    Content   string json:"content"
    Timestamp string json:"timestamp"
}