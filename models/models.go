package models

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
