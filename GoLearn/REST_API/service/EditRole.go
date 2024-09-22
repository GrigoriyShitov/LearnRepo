package service

import (
	"RestApi/storage"
	"context"
)

func EditRoleService(ctx context.Context, whoChange uint, idUser uint, role string) error {
	err := storage.AdminValid(ctx, whoChange)
	if err != nil {
		return err
	}

	storage.EditRole(ctx, idUser, role)

	return nil

}
