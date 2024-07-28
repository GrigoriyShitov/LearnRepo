package service

import (
	"RestApi/storage"
	"context"
	"errors"
)

func GetWalletsInfo(ctx context.Context, id uint) ([]storage.WalletsInfo, error) {

	returnWallets, temp := []storage.WalletsInfo{}, storage.WalletsInfo{}

	for _, wallet := range storage.Wallets {
		if wallet.UserId == id {
			temp.WalletInfo = wallet
			temp.Balance, _ = CalculateBalance(ctx, wallet.WalletId)
			returnWallets = append(returnWallets, temp)
		}
	}
	if len(returnWallets) == 0 {
		return returnWallets, errors.New("no wallets")
	}
	return returnWallets, nil
}
