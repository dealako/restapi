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
	statsd "gopkg.in/alexcesaro/statsd.v2"
)

// RootCmd to kick things off
var RootCmd = &cobra.Command{
	Use:   "restapi",
	Short: "REST API Example",
	Long:  utils.Art(),
}

var httpPort int64 = 8000
var statsHost = "localhost"
var statsPort int64 = 8125
var stats statsd.Client

func init() {
	RootCmd.Run = create

	RootCmd.Flags().Int64VarP(&httpPort, "http-port", "p", 8000, "The HTTP port")
	RootCmd.Flags().StringVarP(&statsHost, "stats-host", "", "localhost", "The stats host to connect to")
	RootCmd.Flags().Int64VarP(&statsPort, "stats-port", "", 8125, "The stats port to connect to")
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
	log.Infof("Statsd host           : %s", statsHost)
	log.Infof("Statsd port           : %d", statsPort)

	stats, err := statsd.New(statsd.Address(fmt.Sprintf("%s:%d", statsHost, statsPort))) // Connect to the UDP host:port
	if err != nil {
		// If nothing is listening on the target port, an error is returned and
		// the returned client does nothing but is still usable. So we can
		// just log the error and go on.
		log.Print(err)
	}
	defer stats.Close()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r))

}

// Init books variable as a slice Books struct
var books []models.Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Increment our stats
	stats.Increment("restapi.book.get.counter")

	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Warn("Error while encoding response: %v", err)
	}
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	// Increment our stats
	stats.Increment("restapi.book.getId.counter")

	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for _, item := range books {
		if item.ID == params["id"] {
			// Write the book back to the client and return
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				log.Warn("Error while encoding response: %v", err)
			}
			return
		}
	}

	// Fell through - write an empty book back to the client and return
	err := json.NewEncoder(w).Encode(&models.Book{})
	if err != nil {
		log.Warn("Error while encoding response: %v", err)
	}
}

// Create Books
func createBook(w http.ResponseWriter, r *http.Request) {
	// Increment our stats
	stats.Increment("restapi.book.create.counter")

	// Set the response type
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Warn("Error while decoding response: %v", err)
	}
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)

	// Write the new book back to the client and return
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Warn("Error while encoding response: %v", err)
	}
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	// Increment our stats
	stats.Increment("restapi.book.update.counter")

	// Set the response type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Look through the books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book models.Book
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Warn("Error while decoding response: %v", err)
			}
			book.ID = params["id"]
			books = append(books, book)

			// Write the updated book back to the client and return
			err = json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Warn("Error while encoding response: %v", err)
			}
			return
		}
	}

	// Return the current slice of books
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Warn("Error while encoding response: %v", err)
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Increment our stats
	stats.Increment("restapi.book.delete.counter")

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
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Warn("Error while encoding response: %v", err)
	}
}

// initMockData creates some sample data and fills the books slice
func initMockData() {

	// Generate Mock Data
	books = append(books, models.GenerateSampleBookRecord())
	books = append(books, models.GenerateSampleBookRecord())
	books = append(books, models.GenerateSampleBookRecord())
	books = append(books, models.GenerateSampleBookRecord())
}

// Execute runs the initial entrypoint to the command line application
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error running tre: %v", err)
	}
}
