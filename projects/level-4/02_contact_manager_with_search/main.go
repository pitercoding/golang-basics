package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	contacts := []Contact{}

	for {
		fmt.Println("\n=== Contact Manager ===")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Remove Contact")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			fmt.Println("\n--- New Contact Details ---")
			name := readString(scanner, "Name: ")
			phone := readString(scanner, "Phone: ")
			email := readString(scanner, "Email: ")

			addContact(&contacts, name, phone, email)

			fmt.Println("Contact added successfully.")
		case 2:
			fmt.Println("\n--- Contacts ---")
			listContacts(contacts)

		case 3:
			name := readString(scanner, "Search by name: ")

			result := searchContact(contacts, name)

			if result != nil {
				fmt.Println("\nContact Found:")
				fmt.Printf("Name: %s | Phone: %s | Email: %s\n",
					result.Name,
					result.Phone,
					result.Email)
			} else {
				fmt.Println("Contact not found.")
			}

		case 4:
			name := readString(scanner, "Name to remove: ")

			removeContact(&contacts, name)

		case 0:
			fmt.Println("\nGoodbye! Closing system...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}

}

func addContact(
	contacts *[]Contact,
	name, phone, email string,
) {
	contact := Contact{
		Name:  name,
		Phone: phone,
		Email: email,
	}

	*contacts = append(*contacts, contact)
}

func listContacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("No contacts found!")
		return
	}

	for i, contact := range contacts {
		fmt.Printf("%d. Name: %s | Phone: %s | Email: %s\n",
			i+1,
			contact.Name,
			contact.Phone,
			contact.Email)
	}
}

func searchContact(contacts []Contact, name string) *Contact {
	for i := range contacts {
		if strings.EqualFold(contacts[i].Name, name) {
			return &contacts[i]
		}
	}
	return nil
}

func removeContact(contacts *[]Contact, name string) {
	for i, contact := range *contacts {
		if strings.EqualFold(contact.Name, name) {
			*contacts = append((*contacts)[:i], (*contacts)[i+1:]...)
			fmt.Println("Contact removed successfully!")
			return
		}
	}

	fmt.Println("Contact not found.")
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("\nInvalid number. Try again.")
	}
}

func readString(scanner *bufio.Scanner, message string) string {
	for {
		fmt.Print(message)

		scanner.Scan()

		value := strings.TrimSpace(scanner.Text())

		if value != "" {
			return value
		}

		fmt.Println("Value cannot be empty! Try again.")
	}
}
