package service

import (
	"RestApi/storage"
	"context"
)

func CreateNewUser(ctx context.Context, id uint, name string) error {
	err := storage.CheckUserExist(ctx, id)

	if err != nil {
		err = storage.CreateNewUser(ctx, id, name)
	}
	return err
}
