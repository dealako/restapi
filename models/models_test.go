package models

import (
	"testing"
)

// TestGetRandomStreetName a test for the getRandomStreetName function
func TestGetRandomStreetName(t *testing.T) {
	street := getRandomStreetName()
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
