package handlers

import (
	"026_rest_api/internal/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /users", h.addUser)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.User{})
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserId: uuid.New()})
}