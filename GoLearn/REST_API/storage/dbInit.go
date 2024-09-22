package storage

import (
	secret "RestApi/.env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func dbInit() {
	var err error
	db, err = gorm.Open(postgres.Open(secret.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("DB connected")
	}
	db = AutomigrateDB(db)
}
