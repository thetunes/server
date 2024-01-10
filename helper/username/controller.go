package helper

import (
	"api/model"
	"errors"

	"gorm.io/gorm"
)

// isUsernameTaken checks if the given username is already registered.
func IsUsernameTaken(db *gorm.DB, username string) bool {
	var existingUser model.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		// Check if the error is due to the record not being found
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}
	// User with the same username already exists
	return true
}
