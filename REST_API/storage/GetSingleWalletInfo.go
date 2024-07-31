package storage

import (
	"context"
	"errors"
)

func GetSingleWalletInfo(ctx context.Context, WalletId uint) (*Wallet, error) { // to do ([]storage.Wallet, error)

	for _, wallet := range Wallets {
		if wallet.Id == WalletId {
			return &wallet, nil
		}
	}
	err := errors.New("wallet not found")
	return &Wallet{}, err

}
