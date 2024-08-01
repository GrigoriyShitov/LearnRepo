package storage

import (
	"context"
	"errors"
)

func WalletValid(ctx context.Context, userID uint, walletID uint) error {
	var (
		wallet Wallet
		user   User
	)

	result := db.First(&wallet, walletID)
	if result.Error != nil {
		return result.Error
	}
	result = db.First(&user, userID)
	if result.Error != nil {
		return result.Error
	}
	if wallet.UserId == userID {
		return nil
	}
	err := errors.New("no access")
	return err

}
