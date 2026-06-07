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

func (u User) Login(username, password string) bool {
	return u.Username == username && u.Password == password
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	user := User{
		Username: "admin",
		Password: "123456",
	}

	fmt.Println("\n=== Login System ===")

	fmt.Print("Username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())


	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	if user.Login(username, password) {
		fmt.Println("\nLogin successful!")
	} else {
		fmt.Println("\nInvalid username or password.")
	}

}
