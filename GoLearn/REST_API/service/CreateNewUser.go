package service

import (
	"RestApi/storage"
	"context"
)

func CreateNewUser(ctx context.Context, name string) error {

	err := storage.CreateNewUser(ctx, name)
	return err
}
