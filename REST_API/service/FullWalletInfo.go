package service

import (
	"RestApi/storage"
	"context"
)

func FullWalletInfo(ctx context.Context, WalletId uint) ([]storage.WalletsInfo, error) {
	wallets, err := storage.GetWalletsInfo(ctx, WalletId)
	if err != nil {
		return nil, err
	}
	returnValue := make([]storage.WalletsInfo, 0, len(wallets))

	for _, wallet := range wallets {
		balance, err := WalletBalanceInfo(ctx, wallet.Id)
		if err != nil {
			return nil, err
		}
		returnValue = append(returnValue, storage.WalletsInfo{
			WalletInfo: wallet,
			Balance:    balance,
		})
	}
	return returnValue, nil
}
