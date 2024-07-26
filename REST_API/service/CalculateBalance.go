package service

import (
	"RestApi/storage"
	"context"
)

type WalletInfo struct {
	Wallet     *storage.Wallet
	Operations []storage.Operation
	Balance    float64 `json:"balance"`
}

func CalculateBalance(ctx context.Context, WalletId uint) (float64, error) {
	var (
		WalletInfo WalletInfo
		err        error
	)
	WalletInfo.Wallet, err = storage.GetSingleWalletInfo(ctx, WalletId)
	WalletInfo.Operations, err = storage.GetOperationOnWallet(ctx, WalletId)
	if err != nil {
		return 0, err
	}
	WalletInfo.Balance = 0
	for _, operation := range WalletInfo.Operations {
		if operation.OperationType == "deposit" {
			WalletInfo.Balance += operation.Amount
		} else if operation.OperationType == "withdraw" {
			WalletInfo.Balance -= operation.Amount
		}
	}
	return WalletInfo.Balance, nil
}
