package service

import (
	"RestApi/storage"
	"context"
)

func NewOperationWithWallet(ctx context.Context, AddOperation *storage.NewOperation) error {

	err := storage.WalletValid(ctx, AddOperation.IdUser, AddOperation.IdWallet)
	if err != nil {
		return err
	}

	err = storage.AddNewOperation(AddOperation)
	if err != nil {
		return err
	}

	return nil
}
