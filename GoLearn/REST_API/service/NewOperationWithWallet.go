package service

import (
	"RestApi/storage"
	"context"
)

func NewOperationWithWallet(ctx context.Context, userID uint, WalletNum uint, Type string, Amount float64, Category string) ([]byte, error) {

	// err := storage.WalletValid(ctx, idUser, idWallet)
	// if err != nil {
	// 	return nil, err
	// }

	err := storage.AddNewOperation(userID, WalletNum, Type, Amount, Category)
	if err != nil {
		return nil, err
	}
	data := []byte("Operation added")
	return data, nil
}
