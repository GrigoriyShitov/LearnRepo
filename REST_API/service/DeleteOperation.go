package service

import (
	"RestApi/storage"
	"context"
)

func DeleteOperation(ctx context.Context, idUser uint, idWallet uint, idOperation uint) error {

	err := storage.WalletValid(ctx, idUser, idWallet)
	if err != nil {

		return err
	}
	err = storage.DeleteOperation(idWallet, idOperation)
	if err != nil {
		return err
	}
	return nil
}
