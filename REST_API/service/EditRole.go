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
	var isAdmin bool
	if role == "admin" {
		isAdmin = true
	}
	storage.EditRole(ctx, idUser, isAdmin)

	return nil

}
