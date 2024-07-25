package methods

import (
	"RestApi/storage"
	"context"
	"errors"
)

func GetWalletsInfo(ctx context.Context, id uint32) ([]storage.Wallet, error) { // to do ([]storage.Wallet, error)

	returnWallets := []storage.Wallet{}

	for _, wallet := range storage.Wallets {
		if wallet.UserId == uint32(id) {
			returnWallets = append(returnWallets, wallet)
		}
	}
	if len(returnWallets) == 0 {
		return returnWallets, errors.New("no wallets")
	}
	return returnWallets, nil
}
