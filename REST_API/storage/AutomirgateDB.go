package storage

import (
	"log"

	"gorm.io/gorm"
)

func AutomigrateDB(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	// db.Create(&User{
	// 	Username: "Grisha",
	// 	Role:     "user",
	// })
	db.AutoMigrate(&Wallet{})
	if err != nil {
		log.Fatal(err)
	}
	// db.Create(&Wallet{
	// 	UserId: 1,
	// })
	// db.Create(&Wallet{
	// 	UserId: 1,
	// })
	// db.Create(&Wallet{
	// 	UserId: 2,
	// })
	// db.Create(&Wallet{
	// 	UserId: 2,
	// })
	db.AutoMigrate(&Operation{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
