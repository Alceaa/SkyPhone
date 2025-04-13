package models

type ContactResponse struct {
	ContactID int    `db:"contactid" json:"contactid"`
	AddBy     string `json:"addBy"`
}
