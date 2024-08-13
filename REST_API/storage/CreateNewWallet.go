package storage

import "context"

func CreateNewWallet(ctx context.Context, id uint) error {

	var wallet Wallet
	wallet.UserId = id

	if err := db.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}
