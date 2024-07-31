package storage

import (
	"context"
	"errors"
)

func WalletValid(ctx context.Context, userID uint, walletID uint) error {
	for _, wallet := range Wallets {
		if wallet.Id == walletID && wallet.UserId == userID {
			return nil
		}
	}
	err := errors.New("it's not your wallet")
	return err

}
