package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string
	Password string
}

type TokenResponse struct {
	Token string
}

var users []User
var nextID = 1

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile", profileHandler)

	fmt.Println("\n=== WELCOME TO USER AUTHENTICATION PAGE ===")
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Authentication API")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Name, Email and Password cannot be empty", http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == user.Email {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

	}

	user.ID = nextID
	nextID++

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")

	response := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.Email == req.Email && u.Password == req.Password {

			token := fmt.Sprintf("TOKEN-%d", u.ID)
			users[i].Token = token

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TokenResponse{
				Token: token,
			})
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)

	for _, u := range users {
		if u.Token == token {

			w.Header().Set("Content-Type", "application/json")

			response := UserResponse{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			}

			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
}
