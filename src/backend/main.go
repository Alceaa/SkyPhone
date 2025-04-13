package main

import (
	"log"
	"net/http"

	"github.com/Alceaa/SkyPhone/db"
	"github.com/Alceaa/SkyPhone/routes"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

var DB *sqlx.DB

func main() {
	DB = db.DB
	r := routes.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
	})
	handler := cors.Default().Handler(r)
	handler = c.Handler(handler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
