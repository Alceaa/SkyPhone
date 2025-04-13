package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Name     string `db:"name" json:"name"`
	Surname  string `db:"surname" json:"surname"`
}
