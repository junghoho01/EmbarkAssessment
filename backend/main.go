package main // So go will run look for this main package to start execution (Executable Go Program)

import ( // To include external package and libraries
	"encoding/json" // Convert go data structure to JSON format and vice versa
	"fmt"
	"net/http" // Provide HTTP client and server implementations GET POST requests
	"strconv"
	"sync" // To handle multiple request safely

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux" // Allow to define routes
	// Import the handlers package for CORS
)

type Book struct { // Define data structure of book
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Image       string `json:"image"`
	Published   string `json:"published"`
	Publisher   string `json:"publisher"`
}

var (
	books  = make(map[int]Book) // Store all book with an int as the key and book as the details (Just stored in RAM)
	nextID = 1                  // Keeps track of the next available ID
	mu     sync.Mutex           // To avoid corruption of data where trying to read and write simultaneously
)

var defaultBooks = []Book{
	{
		ID:          1,
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan",
		Genre:       "Programming",
		Description: "An introduction to Go.",
		ISBN:        "978-0134190440",
		Image:       "http://example.com/go-book.jpg",
		Published:   "2015",
		Publisher:   "Addison-Wesley",
	},
	{
		ID:          2,
		Title:       "Clean Code",
		Author:      "Robert C. Martin",
		Genre:       "Programming",
		Description: "A Handbook of Agile Software Craftsmanship.",
		ISBN:        "978-0132350884",
		Image:       "http://example.com/clean-code.jpg",
		Published:   "2008",
		Publisher:   "Prentice Hall",
	},
	{
		ID:          3,
		Title:       "The Pragmatic Programmer",
		Author:      "Andrew Hunt, David Thomas",
		Genre:       "Programming",
		Description: "Your Journey to Mastery.",
		ISBN:        "978-0135957059",
		Image:       "http://example.com/pragmatic-programmer.jpg",
		Published:   "1999",
		Publisher:   "Addison-Wesley",
	},
	{
		ID:          4,
		Title:       "The Clean Coder",
		Author:      "Robert C. Martin",
		Genre:       "Programming",
		Description: "A Code of Conduct for Professional Programmers.",
		ISBN:        "978-0136083238",
		Image:       "http://example.com/clean-coder.jpg",
		Published:   "2011",
		Publisher:   "Prentice Hall",
	},
	{
		ID:          5,
		Title:       "Refactoring: Improving the Design of Existing Code",
		Author:      "Martin Fowler",
		Genre:       "Programming",
		Description: "A detailed guide on how to refactor code.",
		ISBN:        "978-0201485677",
		Image:       "http://example.com/refactoring.jpg",
		Published:   "1999",
		Publisher:   "Addison-Wesley",
	},
	{
		ID:          6,
		Title:       "Code Complete",
		Author:      "Steve McConnell",
		Genre:       "Software Development",
		Description: "A practical guide to software construction.",
		ISBN:        "978-0735619678",
		Image:       "http://example.com/code-complete.jpg",
		Published:   "2004",
		Publisher:   "Microsoft Press",
	},
	{
		ID:          7,
		Title:       "You Don't Know JS",
		Author:      "Kyle Simpson",
		Genre:       "Programming",
		Description: "An in-depth look at JavaScript.",
		ISBN:        "978-1491903995",
		Image:       "http://example.com/you-dont-know-js.jpg",
		Published:   "2015",
		Publisher:   "O'Reilly Media",
	},
}

type Response struct {
	Message string `json:"message"`
	Book    Book   `json:"book"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

func createBook(w http.ResponseWriter, r *http.Request) { // Post request to create new book
	mu.Lock()                                                     // Can only executed by one goroutine at a time
	defer mu.Unlock()                                             // Unlock is guranteed after the function finished
	var book Book                                                 // Create a book of type of empty book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil { // decode the json data and store in book var
		http.Error(w, err.Error(), http.StatusBadRequest) // if has error when decoding, send an error back to the client
		return
	}
	book.ID = nextID // success decode, assign id to the book
	nextID++
	books[book.ID] = book // store the new book to the book in RAM
	// Create the response object with message and book
	response := Response{
		Message: "New Book Added Successfully",
		Book:    book,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Encode the response with the message and book, send the new book back as a JSON object confirm it is created
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extract the book ID from the URL path
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr) // Convert the ID from string to int
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	mu.Lock() // Lock to ensure thread-safe access to the `books` map
	defer mu.Unlock()

	if _, ok := books[id]; !ok { // Check if the book with the given ID exists
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id) // Delete the book from the `books` map

	// Send a success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the URL path
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the ID from string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	mu.Lock() // Lock the map for thread-safe access
	defer mu.Unlock()

	// Check if the book exists
	if _, ok := books[id]; !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Parse the updated book details from the request body
	var updatedBook Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retain the original book ID
	updatedBook.ID = id

	// Update the book in the `books` map
	books[id] = updatedBook

	// Send a success message with the updated book
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book updated successfully",
		"book":    updatedBook,
	})
}

func resetBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Clear the current books map
	books = make(map[int]Book)

	// Reset the `nextID` counter
	nextID = 1

	// Add the default set of books to the `books` map
	for _, book := range defaultBooks {
		book.ID = nextID
		books[nextID] = book
		nextID++
	}

	// Send a success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book list reset to its original state"})
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extract the book ID from the URL path
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr) // Convert the ID from string to int
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	mu.Lock() // Lock to ensure thread-safe access to the `books` map
	defer mu.Unlock()

	book, ok := books[id] // Check if the book with the given ID exists
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Send the book details as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()

	// Serve static files from the Vue.js app
	http.Handle("/", http.FileServer(http.Dir("./book-management/dist")))

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/books/reset", resetBooks).Methods("POST")

	// Enable CORS for all routes
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins (for development)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	// Print the message to indicate that the server is running
	fmt.Println("Server is running on port 8080...")

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", corsHandler)
}
