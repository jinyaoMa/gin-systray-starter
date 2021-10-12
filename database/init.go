package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connect error %v\n", err)
	}
	if DB.Error != nil {
		log.Fatalf("database error %v\n", DB.Error)
	}
}
