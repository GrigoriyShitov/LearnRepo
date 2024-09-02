package service

import (
	"RestApi/storage"
	"context"
)

func UserWalletInfo(ctx context.Context, id uint) ([]byte, error) {

	Username, Role, err := storage.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	WalletsList, err := FullWalletInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	data, err := UserInfoToOut(ctx, Username, Role, WalletsList)
	if err != nil {
		return nil, err
	}
	return data, err
}
