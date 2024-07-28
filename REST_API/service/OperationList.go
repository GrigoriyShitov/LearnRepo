package service

import (
	"RestApi/storage"
	"context"
)

type OperationToOut struct {
	OperationType     string
	Amount            float64
	OperationCategory string
}

func OperationList(ctx context.Context, id uint, idWallet uint) ([]OperationToOut, error) {
	err := storage.WalletValid(ctx, id, idWallet)
	if err != nil {
		return nil, err
	}
	ListOfOperations, _ := storage.GetOpList(ctx, idWallet)

	Output := ListToOut(ctx, ListOfOperations)

	return Output, nil
}
