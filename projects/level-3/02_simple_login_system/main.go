package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
}

func main() {
	users, err := loadUsers("users.txt")
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	var username string
	var password string

	fmt.Println("\n=== Simple Login System ===")

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if authenticate(users, username, password) {
		fmt.Println("\nLogin successful!")
	} else {
		fmt.Println("\nInvalid username or password.")
	}
}

func loadUsers(filePath string) ([]User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			continue
		}

		user := User{
			Username: parts[0],
			Password: parts[1],
		}

		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func authenticate(
	users []User,
	username string,
	password string,
) bool {
	for _, user := range users {
		if user.Username == username &&
			user.Password == password {
			return true
		}
	}

	return false
}
