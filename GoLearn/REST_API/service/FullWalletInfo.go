package service

import (
	"RestApi/storage"
	"context"
)

type WalletsInfo struct {
	WalletId uint    `json:"wallet_id"`
	OwnerId  uint    `json:"owner_id"`
	Balance  float64 `json:"balance"`
}

func FullWalletInfo(ctx context.Context, WalletId uint) ([]WalletsInfo, error) {
	wallets, err := storage.GetWalletsInfo(ctx, WalletId)
	if err != nil {
		return nil, err
	}
	returnValue := make([]WalletsInfo, 0, len(wallets))

	for _, wallet := range wallets {
		balance, err := WalletBalanceInfo(ctx, wallet.Id)
		if err != nil {
			return nil, err
		}
		returnValue = append(returnValue, WalletsInfo{
			WalletId: wallet.Id,
			OwnerId:  wallet.UserId,
			Balance:  balance,
		})
	}
	return returnValue, nil
}
