package models

type Chat struct {
	ChatID string `db:"id" json:"chatid"`
	Name   string `db:"name" json:"name"`
	Count  string `db:"count" json:"count"`
}
