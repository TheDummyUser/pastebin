package db

import (
	"fmt"

	"github.com/thedummyuser/pastebin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("pastbin.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to the database", err)
		return nil, err
	}

	db.AutoMigrate(&models.Paste{})

	fmt.Println("auto migration successful")
	return db, nil
}
