package repository

import (
	"errors"

	"api/database"
	"api/model"

	"gorm.io/gorm"
)

func FindByCredentials(username, password string) (*model.User, error) {
	var user model.User
	// Here you would query your MySQL database for the user with the given username and password
	result := database.DB.Db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func AdminFindByCredentials(username, password string) (*model.Admin, error) {
	var admin model.Admin
	// Here you would query your MySQL database for the Admin with the given username and password
	result := database.DB.Db.Where("username = ? AND password = ?", username, password).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Admin not found")
		}
		return nil, result.Error
	}
	return &admin, nil
}
