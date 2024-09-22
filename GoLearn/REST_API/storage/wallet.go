package storage

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Id     uint `json:"wallet_id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
}