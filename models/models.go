package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

// Address Struct
type Address struct {
	AddressLine1 string `json:"address1"`
	AddressLine2 string `json:"address2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          int32  `json:"zip"`
}

var recordCount = 0

// GenerateSampleBookRecord generates and returns a sample Book record
func GenerateSampleBookRecord() Book {
	// Increment our record count
	recordCount++

	// Mock Data - Book
	book := Book{
		ID:    strconv.Itoa(recordCount),
		Isbn:  strconv.Itoa(rand.Intn(100000)),
		Title: fmt.Sprintf("Book %d", recordCount),
	}

	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	address := Address{
		AddressLine1: fmt.Sprintf("%d Main Street", rand.Intn(1000)),
		AddressLine2: "",
		City:         "San Diego",
		State:        "CA",
		Zip:          92127,
	}

	author.Address = &address
	book.Author = &author

	return book
}
