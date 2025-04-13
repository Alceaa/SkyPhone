package models

type MessageResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
