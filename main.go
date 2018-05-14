package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/dealako/restapi/models"
	"github.com/gorilla/mux"
)

// Build and version variables defined and set during the build process
var (
	// Name the application name
	name string
	// Version the application version
	version string
	// Build/Commit the application build number
	commit string
	// Build date
	buildDate string
)

// Init books variable as a slice Books struct
var books []models.Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Look through the books and find the id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

// Create Books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)

			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func initMockData() {

	// Mock Data
	book1 := models.Book{
		ID:    "1",
		Isbn:  "12345",
		Title: "Book One",
	}
	author1 := models.Author{
		Firstname: "John",
		Lastname:  "Doe",
	}
	address1 := models.Address{
		AddressLine1: "123 Main Street",
		AddressLine2: "",
		City:         "San Diego",
		State:        "CA",
		Zip:          92127,
	}

	author1.Address = &address1
	book1.Author = &author1
	books = append(books, book1)

	book2 := models.Book{
		ID:    "2",
		Isbn:  "78910",
		Title: "Book Two",
	}
	author2 := models.Author{
		Firstname: "Sam",
		Lastname:  "Smith",
	}
	address2 := models.Address{
		AddressLine1: "456 Front Street",
		AddressLine2: "",
		City:         "San Diego",
		State:        "CA",
		Zip:          92127,
	}

	author2.Address = &address2
	book2.Author = &author2
	books = append(books, book2)
}

func main() {
	// Show the version and build info
	log.Infof("Name                  : %s", name)
	log.Infof("Version               : %s", version)
	log.Infof("Git commit hash       : %s", commit)
	log.Infof("Build date            : %s", buildDate)
	log.Infof("Golang OS             : %s", runtime.GOOS)
	log.Infof("Golang Arch           : %s", runtime.GOARCH)

	// Init Router
	r := mux.NewRouter()

	// Initialize some Mock Data
	initMockData()

	// Route Handlers / Endponts
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
