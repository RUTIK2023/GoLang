// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Book interface
type Book interface {
	GetDetails() string
	IsAvailable() bool
}

// PhysicalBook struct
type PhysicalBook struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// EBook struct
type EBook struct {
	Title     string
	Author    string
	ISBN      string
	FileSize  int
	Available bool
}

// GetDetails method for PhysicalBook
func (p PhysicalBook) GetDetails() string {
	availability := "Yes"
	if !p.Available {
		availability = "No"
	}
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, Available: %s", p.Title, p.Author, p.ISBN, availability)
}

// IsAvailable method for PhysicalBook
func (p PhysicalBook) IsAvailable() bool {
	return p.Available
}

// GetDetails method for EBook
func (e EBook) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, File Size: %d MB, Available: %v", e.Title, e.Author, e.ISBN, e.FileSize, e.Available)
}

// IsAvailable method for EBook
func (e EBook) IsAvailable() bool {
	return e.Available
}

// Library struct
type Library struct {
	Books []Book
}

// AddBook method
func (lib *Library) AddBook(book Book) {
	lib.Books = append(lib.Books, book)
	fmt.Printf("Book added: %s\n", book.GetDetails())
}

// RemoveBook method
func (lib *Library) RemoveBook(isbn string) {
	found := false
	for i, book := range lib.Books {
		if b, ok := book.(PhysicalBook); ok && b.ISBN == isbn {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Printf("Physical book removed: %s\n", b.Title)
			found = true
			break
		} else if e, ok := book.(EBook); ok && e.ISBN == isbn {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Printf("EBook removed: %s\n", e.Title)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Book with ISBN %s not found.\n", isbn)
	}
}

// ListBooks method
func (lib *Library) ListBooks() {
	fmt.Println("Library books:")
	for _, book := range lib.Books {
		fmt.Println(book.GetDetails())
	}
}

// SearchByTitle method
func (lib *Library) SearchByTitle(title string) {
	found := false
	fmt.Println("Search results:")
	for _, book := range lib.Books {
		if strings.Contains(strings.ToLower(book.GetDetails()), strings.ToLower(title)) {
			fmt.Println(book.GetDetails())
			found = true
		}
	}
	if !found {
		fmt.Printf("No books found with title containing '%s'.\n", title)
	}
}

// Main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	library := Library{}
	running := true

	for running {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Physical Book")
		fmt.Println("2. Add an EBook")
		fmt.Println("3. Remove a Book by ISBN")
		fmt.Println("4. Search for a Book by Title")
		fmt.Println("5. List All Books")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// Add a Physical Book
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			fmt.Print("Enter ISBN: ")
			scanner.Scan()
			isbn := scanner.Text()

			library.AddBook(PhysicalBook{Title: title, Author: author, ISBN: isbn, Available: true})

		case "2":
			// Add an EBook
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			fmt.Print("Enter ISBN: ")
			scanner.Scan()
			isbn := scanner.Text()

			fmt.Print("Enter File Size (MB): ")
			scanner.Scan()
			var fileSize int
			fmt.Sscan(scanner.Text(), &fileSize)

			library.AddBook(EBook{Title: title, Author: author, ISBN: isbn, FileSize: fileSize, Available: true})

		case "3":
			// Remove a Book by ISBN
			fmt.Print("Enter ISBN of the book to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			library.RemoveBook(isbn)

		case "4":
			// Search for a Book by Title
			fmt.Print("Enter Title to search: ")
			scanner.Scan()
			title := scanner.Text()
			library.SearchByTitle(title)

		case "5":
			// List All Books
			library.ListBooks()

		case "6":
			// Exit
			running = false
			fmt.Println("Exiting the system. Goodbye!")

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
