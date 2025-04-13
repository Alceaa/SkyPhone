package models

type Message struct {
	ID        int    `json:"id"`
	ChatID    string `json:"chat_id"`
	SenderID  int    `json:"sender_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
