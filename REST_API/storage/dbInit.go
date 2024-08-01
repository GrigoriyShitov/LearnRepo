package storage

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func dbInit() {
	dsn := "host=localhost user=grish password=1129 dbname=restapi port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("DB connected")
	}
	db = AutomigrateDB(db)
}
