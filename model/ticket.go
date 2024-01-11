package model

import (
	"gorm.io/gorm"
)

// Define data struct required for User
type Ticket struct {
	gorm.Model
	ID          string   `json:"id"`
	ArtistID    string   `json:"artistid" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
	Location    string   `json:"location" validate:"required"`
	Date        string   `json:"date" validate:"required"`
	Description string   `json:"description"`
	Likes       int      `json:"likes"`
	Likers      []string `json:"likers" gorm:"type:uuid[]"` // Add this line
	PromotorID  string   `json:"promotorid"`
	Status      string   `json:"status"`
}

// Users struct
type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}
