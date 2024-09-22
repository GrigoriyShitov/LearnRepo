package service

import (
	"RestApi/storage"
	"context"
)

func CreateNewWallet(ctx context.Context, id uint) error {
	err := storage.CheckUserExist(ctx, id)

	if err != nil {
		return err
	}
	err = storage.CreateNewWallet(ctx, id)
	return err
}
