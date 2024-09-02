package storage

import "gorm.io/gorm"

type Operation struct {
	gorm.Model
	WalletId          uint    `json:"wallet_id"`
	Amount            float64 `json:"amount"`
	OperationType     string  `json:"operation_type"`
	OperationCategory string  `json:"operation_category"`
}

type NewOperation struct {
	IdUser   uint
	IdWallet uint
	Type     string
	Amount   float64
	Category string
}
