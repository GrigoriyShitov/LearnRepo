package service

import (
	"RestApi/storage"
	"context"
)

func UserWalletInfo(ctx context.Context, id uint) (*storage.UserInfo, error) {
	var (
		RetVal storage.UserInfo
		err    error
	)
	RetVal.User, err = storage.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	RetVal.WalletsList, err = FullWalletInfo(ctx, id)
	if err != nil {
		return &RetVal, err
	}
	return &RetVal, nil
}
