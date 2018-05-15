package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/dealako/restapi/models"
	"github.com/dealako/restapi/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// RootCmd to kick things off
var RootCmd = &cobra.Command{
	Use:   "restapi",
	Short: "REST API Example",
	Long:  utils.Art(),
}

var httpPort int64 = 8000

func init() {
	RootCmd.Run = create

	RootCmd.Flags().Int64VarP(&httpPort, "port", "p", 8000, "The HTTP port")
}

func create(cmd *cobra.Command, args []string) {
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

	log.Infof("HTTP server port      : %d", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r))

}

// Init books variable as a slice Books struct
var books []models.Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for _, item := range books {
		if item.ID == params["id"] {
			// Write the book back to the client and return
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Fell through - write an empty book back to the client and return
	json.NewEncoder(w).Encode(&models.Book{})
}

// Create Books
func createBook(w http.ResponseWriter, r *http.Request) {
	// Set the response type
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)

	// Write the new book back to the client and return
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	// Set the response type
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

			// Write the updated book back to the client and return
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// Return the current slice of books
	json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			// Strip out the matching record from the the slice and return
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	// Return the current slice of books
	json.NewEncoder(w).Encode(books)
}

// initMockData creates some sample data and fills the books slice
func initMockData() {

	// Mock Data
	books = append(books, models.GenerateSampleBookRecord())
	books = append(books, models.GenerateSampleBookRecord())
	books = append(books, models.GenerateSampleBookRecord())
}

// Execute runs the initial entrypoint to the tre cli
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error running tre: %v", err)
	}
}
