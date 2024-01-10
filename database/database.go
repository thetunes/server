package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"api/config"
	"api/model"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connection to database
func Connect() {
	p := config.Config("DB_PORT")

	// Parse data returned by config function as string
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running DB Auto Migration for Admin")
	db.AutoMigrate(&model.User{}, &model.Admin{}, &model.Ticket{}, &model.Promotor{})
	DB = Dbinstance{
		Db: db,
	}
}
