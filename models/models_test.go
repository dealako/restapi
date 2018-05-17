package models

import (
	"testing"
)

// TestGetRandomStreetNameNotNil
func TestGetRandomStreetNameNotNil(t *testing.T) {
	street, err := getRandomStreetName()
	if err != nil {
		t.Error(err)
	}

	if street == "" {
		t.Fatal("String name is empty.")
	}
}

// TestGetRandomStreetName a test for the getRandomStreetName function
func TestGetRandomStreetName(t *testing.T) {
	street, err := getRandomStreetName()
	if err != nil {
		t.Error(err)
	}

	var found = false
	for _, s := range streets {
		if s == street {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Street not found: %s in slice: %v", street, streets)
	}
}
