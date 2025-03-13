package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

func main() {
	//r := mux.NewRouter()
	//r.HandleFunc("/ws", handleChat)
	log.Fatal(http.ListenAndServe(":8080", r))
}