package main

import (
	"fmt"
)

func main() {
	library := Library{Books: make(map[string]Book)}
	for {
		fmt.Println("\n1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scan(&author)
			library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case 2:
			var id string
			fmt.Print("Enter Book ID to borrow: ")
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			var id string
			fmt.Print("Enter Book ID to return: ")
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			library.ListBooks()
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("Book added successfully!")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found!")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed!")
		return
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed successfully!")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found!")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book was not borrowed!")
		return
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned successfully!")
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
