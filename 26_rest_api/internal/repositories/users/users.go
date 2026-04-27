package users

import "026_rest_api/internal/models"

type Users struct {
	users []models.User
}

func New() *Users {
	return &Users{users: make([]models.User, 0)}
}

func (u Users) GetAll() []models.User {
	return u.users
}

func (u Users) EmailExists(email string) (bool) {
	for _, v := range u.users {
		if v.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser models.User) {
	u.users = append(u.users, newUser)
}