package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var messages = []Message{
	{ID: 1, Content: "Hello, World!"},
	{ID: 2, Content: "Golang is awesome!"},
	{ID: 3, Content: "3"},
}

func main() {
	http.HandleFunc("/api/messages", handleMessages)
	log.Println("🚀 Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(messages)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
