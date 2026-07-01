package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var messages []Message
var idCounter = 1

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/messages", messagesHandler)
	http.HandleFunc("/messages/", messageByIDHandler)

	fmt.Println("\n=== WELCOME TO REST API CRUD ===")
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to REST API CRUD!")
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)

	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)

		msg.ID = idCounter
		idCounter++

		messages = append(messages, msg)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func messageByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/messages/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1

	for i, msg := range messages {
		if msg.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages[index])

	case http.MethodPut:
		var updated Message
		json.NewDecoder(r.Body).Decode(&updated)

		messages[index].Content = updated.Content

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages[index])

	case http.MethodDelete:
		messages = append(messages[:index], messages[index+1:]...)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
