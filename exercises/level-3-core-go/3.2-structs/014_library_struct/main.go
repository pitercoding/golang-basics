package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Name  string
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books found.")
		return
	}

	fmt.Printf("\nBooks in %s:\n", l.Name)

	for i, book := range l.Books {
		fmt.Printf("%d. %s - %s\n",
			i+1,
			book.Title,
			book.Author,
		)
	}
}

func (l Library) TotalBooks() int {
	return len(l.Books)
}

func main() {
	library := Library{
		Name: "City Library",
	}

	library.AddBook(Book{
		Title:  "Clean Code",
		Author: "Robert Martin",
	})

	library.AddBook(Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
	})

	library.ListBooks()
	fmt.Printf(
		"\nTotal books: %d\n",
		library.TotalBooks(),
	)
}
