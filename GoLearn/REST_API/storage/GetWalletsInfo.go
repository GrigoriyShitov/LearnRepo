package storage

import (
	"context"
	"errors"
)

func GetWalletsInfo(ctx context.Context, id uint) ([]Wallet, error) {
	Wallets := []Wallet{}
	result := db.Find(&Wallets, "user_id = ?", id)
	if len(Wallets) == 0 {
		return Wallets, errors.New("no wallets")
	}

	if result.Error != nil {
		return Wallets, result.Error
	}

	return Wallets, nil
}
