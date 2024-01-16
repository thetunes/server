package model

import (
	"gorm.io/gorm"
)

// Define data struct required for User
type TicketsOrder struct {
	gorm.Model
	ID       string `json:"id"`
	TicketID string `json:"ticketid"`
	UserID   string `json:"userid"`
	Status   string `json:"status"`
}

// Define data struct required for User
type FileInfo struct {
	ID string `json:"id"`
}

// Ticket Orders struct
type TicketsOrders struct {
	TicketsOrders []TicketsOrder `json:"tickets"`
}
