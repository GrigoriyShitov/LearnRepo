package methods

import (
	"RestApi/storage"
	"context"
)

func GetWalletsInfo(ctx context.Context, id uint32) []storage.Wallet { // to do ([]storage.Wallet, error)

	returnWallets := []storage.Wallet{}

	for _, wallet := range storage.Wallets {
		if wallet.UserId == uint32(id) {
			returnWallets = append(returnWallets, wallet)
		}
	}
	return returnWallets
}
