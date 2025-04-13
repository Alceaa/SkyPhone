package models

type Contact struct {
	UserID    int `db:"userid" json:"userid"`
	ContactID int `db:"contactid" json:"contactid"`
}
