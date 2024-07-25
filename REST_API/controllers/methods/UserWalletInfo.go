package methods

import (
	"RestApi/storage"
	"context"
)

func UserWalletInfo(ctx context.Context, id uint32) (*storage.User, []storage.Wallet, error) { // to do (*storage.User, error)

	User, err := GetUserInfo(ctx, uint32(id))
	if err != nil {
		return nil, nil, err
	}
	Wallets, err := GetWalletsInfo(ctx, uint32(id))
	if err != nil {
		return User, Wallets, err
	}
	return User, Wallets, nil
}
