package storage

import (
	"context"
	"errors"
)

func WalletValid(ctx context.Context, userID uint, WalletNum uint) ([]Wallet, error) {
	var (
		user User
	)
	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	Wallets := []Wallet{}
	result = db.Find(&Wallets, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(Wallets) == 0 {
		return nil, errors.New("no wallets")
	}

	return Wallets, nil

}
