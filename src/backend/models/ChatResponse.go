package models

type ChatResponse struct {
	Name      string `json:"name"`
	Users     []int  `json:"users"`
	CreatedBy string `json:"createdBy"`
}
