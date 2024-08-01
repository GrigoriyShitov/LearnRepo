package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Role     string `json:"role"`
	//Wallets  []Wallet `gorm:"foreignKey:UserId"`
}

type UserInfo struct {
	User        *User
	WalletsList []WalletsInfo
}
