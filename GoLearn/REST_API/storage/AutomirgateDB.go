package storage

import (
	"log"

	"gorm.io/gorm"
)

func AutomigrateDB(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(Wallet{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(Operation{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
