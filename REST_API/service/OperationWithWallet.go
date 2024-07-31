package service

import (
	"RestApi/storage"
	"context"
	"errors"
)

func NewOperationWithWallet(ctx context.Context, AddOperation *storage.NewOperation) error {

	for _, wallet := range storage.Wallets {
		if wallet.Id == AddOperation.WalletId && wallet.UserId == AddOperation.UserId {
			err := storage.AddNewOperation(AddOperation)
			if err != nil {
				err = errors.New("cant make operation")
				return err
			}
			break
		}
	}
	return nil
}
