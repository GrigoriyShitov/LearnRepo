package service

import (
	"RestApi/storage"
	"context"
	"errors"

	"gorm.io/gorm"
)

func GetWalletsInfo(ctx context.Context, id uint, db *gorm.DB) ([]storage.Wallet, error) {

	returnWallets := []storage.Wallet{}
	result := db.Find(&returnWallets, "user_id = ?", id)
	// for _, wallet := range storage.Wallets {
	// 	if wallet.UserId == id {
	// 		temp.WalletInfo = wallet
	// 		temp.Balance, _ = CalculateBalance(ctx, wallet.Id)
	// 		returnWallets = append(returnWallets, temp)
	// 	}
	// }

	if result.Error != nil {
		return returnWallets, result.Error
	}
	if len(returnWallets) == 0 {
		return returnWallets, errors.New("no wallets")
	}
	return returnWallets, nil
}
