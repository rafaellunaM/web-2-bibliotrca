package models

type Book struct {
	Bookid       string  `json:"bookid"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Quantity     int     `json:"quantity"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}
