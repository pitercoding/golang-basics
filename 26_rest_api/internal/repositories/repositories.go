package repositories

import (
	"026_rest_api/internal/models"
	"026_rest_api/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}