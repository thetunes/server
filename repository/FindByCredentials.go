package repository

import (
	"errors"

	"api/database"
	"api/model"

	"gorm.io/gorm"
)

func FindByCredentials(username, password string) (*model.User, error) {
	var user model.User
	// Query the MySQL database for the user with the given username and password
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
	// Query the MySQL database for the Admin with the given username and password
	result := database.DB.Db.Where("username = ? AND password = ?", username, password).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		return nil, result.Error
	}
	return &admin, nil
}

func PromotorFindByCredentials(username, password string) (*model.Promotor, error) {
	var promotor model.Promotor
	// Query the MySQL database for the Promotor with the given username and password
	result := database.DB.Db.Where("username = ? AND password = ?", username, password).First(&promotor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("promotor not found")
		}
		return nil, result.Error
	}
	return &promotor, nil
}
