package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var contacts []Contact
	nextID := 1

	for {
		fmt.Println("\n=== Contact Manager ===")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Search Contact")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {

		case 1:
			name := readString(scanner, "Name: ")
			email := readString(scanner, "Email: ")

			addContact(
				&contacts,
				&nextID,
				name,
				email,
			)

			fmt.Println("Contact added successfully.")

		case 2:
			listContacts(contacts)

		case 3:
			id := readInt(scanner, "Enter contact ID: ")

			name := readString(scanner, "New name: ")
			email := readString(scanner, "New email: ")

			if updateContact(
				contacts,
				id,
				name,
				email,
			) {
				fmt.Println("Contact updated.")
			} else {
				fmt.Println("Contact not found.")
			}

		case 4:
			id := readInt(scanner, "Enter contact ID: ")

			if deleteContact(
				&contacts,
				id,
			) {
				fmt.Println("Contact deleted.")
			} else {
				fmt.Println("Contact not found.")
			}

		case 5:
			
			name := readString(scanner, "Enter name: ")

			contact, found := searchContact(contacts, name)

			if found {
				fmt.Printf(
					"ID: %d | Name: %s | Email: %s\n",
					contact.ID,
					contact.Name,
					contact.Email,
				)
			} else {
				fmt.Println("Contact not found.")
			}

		case 0:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addContact(
	contacts *[]Contact,
	nextID *int,
	name string,
	email string,
) {
	contact := Contact{
		ID:    *nextID,
		Name:  name,
		Email: email,
	}

	*contacts = append(*contacts, contact)

	*nextID++
}

func listContacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	for _, contact := range contacts {
		fmt.Printf(
			"ID: %d | Name: %s | Email: %s\n",
			contact.ID,
			contact.Name,
			contact.Email,
		)
	}
}

func updateContact(
	contacts []Contact,
	id int,
	name string,
	email string,
) bool {
	for i := range contacts {
		if contacts[i].ID == id {
			contacts[i].Name = name
			contacts[i].Email = email
			return true
		}
	}

	return false
}

func deleteContact(
	contacts *[]Contact,
	id int,
) bool {
	for i, contact := range *contacts {
		if contact.ID == id {
			*contacts = append(
				(*contacts)[:i],
				(*contacts)[i+1:]...,
			)

			return true
		}
	}

	return false
}

func searchContact(
	contacts []Contact,
	name string,
) (Contact, bool) {

	for _, contact := range contacts {
		if strings.EqualFold(contact.Name, name) {
			return contact, true
		}
	}

	return Contact{}, false
}

func readString(
	scanner *bufio.Scanner,
	message string,
) string {
	for {
		fmt.Print(message)

		scanner.Scan()

		value := strings.TrimSpace(scanner.Text())

		if value != "" {
			return value
		}

		fmt.Println("Value cannot be empty.")
	}
}

func readInt(
	scanner *bufio.Scanner,
	message string,
) int {
	for {
		fmt.Print(message)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid number.")
	}
}
