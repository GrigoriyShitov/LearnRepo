package db

import (
	"RestApi/storage"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=grish password=1129 dbname=restapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("DB connected")
	}
	db = storage.AutomigrateDB(db)
	return db
}
