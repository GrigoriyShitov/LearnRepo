package service

import (
	"RestApi/storage"
	"context"

	"gorm.io/gorm"
)

func UserWalletInfo(ctx context.Context, id uint, db *gorm.DB) (*storage.UserInfo, error) {
	var (
		RetVal storage.UserInfo
		err    error
	)
	RetVal.User, err = storage.GetUserInfo(ctx, id, db)
	if err != nil {
		return nil, err
	}
	RetVal.WalletsList, err = GetWalletsInfo(ctx, id, db)
	if err != nil {
		return &RetVal, err
	}
	return &RetVal, nil
}
