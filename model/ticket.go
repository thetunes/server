package model

import (
	"gorm.io/gorm"
)

// Define data struct required for User
type Ticket struct {
	gorm.Model
	ID          string `json:"id"`
	ArtistID    string `json:"artistid" gorm:"column:artistid"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Sold        string `json:"sold"`
}

// Users struct
type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}
